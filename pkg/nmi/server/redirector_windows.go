package server

import (
	"encoding/json"
	"fmt"

	"github.com/Microsoft/hcnproxy/pkg/client"
	msg "github.com/Microsoft/hcnproxy/pkg/types"
	v1 "github.com/Microsoft/hcsshim"
	"github.com/Microsoft/hcsshim/hcn"
	"k8s.io/klog/v2"
)

type WindowsRedirector struct {
	MetadataIP   string
	MetadataPort string
	RedirectIP   string
	RedirectPort string
}

func NewRedirector(metadataIP, metadatPort, redirectIP, redirectPort string) *WindowsRedirector {
	return &WindowsRedirector{
		MetadataIP:   metadataIP,
		MetadataPort: metadatPort,
		RedirectIP:   redirectIP,
		RedirectPort: redirectPort,
	}
}

func (wr *WindowsRedirector) AddRedirectRules() error {
	klog.Infof("Enumerating all networks")
	request := msg.HNSRequest{
		Entity:    msg.Network,
		Operation: msg.Enumerate,
		Request:   nil,
	}
	response, err := callHcnProxyAgentInternal(request)
	if err != nil {
		klog.Errorf("failed to enumerate, err: %+v", err)
		return err
	}
	var networks []v1.HNSNetwork
	err = json.Unmarshal(response, &networks)
	if err != nil {
		klog.Errorf("failed to unmarshal networks, err: %+v", err)
		return err
	}
	err = json.Unmarshal(response, &networks)
	if err != nil {
		klog.Errorf("failed to unmarshal networks, err: %+v", err)
		return err
	}
	for _, nw := range networks {
		if nw.Name == "azure" {
			proxy := hcn.L4ProxyPolicySetting{
				IP:          wr.MetadataIP,
				Port:        wr.MetadataPort,
				Protocol:    hcn.ProtocolTypeTCP,
				Destination: fmt.Sprintf("%s:%s", wr.RedirectIP, wr.RedirectPort),
				OutboundNAT: true,
			}
			setting, err := json.Marshal(proxy)
			if err != nil {
				klog.Errorf("failed to marshal proxy, err: %+v", err)
				return err
			}
			policy := hcn.NetworkPolicy{
				Type:     hcn.NetworkL4Proxy,
				Settings: setting,
			}
			networkRequest := hcn.PolicyNetworkRequest{
				Policies: []hcn.NetworkPolicy{policy},
			}
			policySetting, err := json.Marshal(networkRequest)
			if err != nil {
				klog.Errorf("failed to marshal policy setting, err: %+v", err)
				return err
			}
			requestSetting := hcn.ModifyNetworkSettingRequest{
				ResourceType: hcn.NetworkResourceTypePolicy,
				RequestType:  hcn.RequestTypeAdd,
				Settings:     policySetting,
			}
			klog.Infof("adding proxy policy to %s network id: %s", nw.Name, nw.Id)
			clientReq := msg.ModifyNetworkRequest{
				NetworkID:      nw.Id,
				RequestSetting: requestSetting,
			}
			clientR, err := json.Marshal(clientReq)
			if err != nil {
				klog.Errorf("failed to marshal client request, err: %+v", err)
				return err
			}
			clientRequest := msg.HNSRequest{
				Operation: msg.Modify,
				Entity:    msg.Network,
				Request:   clientR,
			}
			response, err = callHcnProxyAgentInternal(clientRequest)
			if err != nil {
				klog.Errorf("failed to modify, err: %+v", err)
				return err
			}
		}
	}
	return nil
}

func (wr *WindowsRedirector) RemoveRedirectRules() error {
	klog.Infof("Enumerating all networks")
	request := msg.HNSRequest{
		Entity:    msg.Network,
		Operation: msg.Enumerate,
		Request:   nil,
	}
	response, err := callHcnProxyAgentInternal(request)
	if err != nil {
		klog.Errorf("failed to enumerate, err: %+v", err)
		return err
	}
	var networks []v1.HNSNetwork
	err = json.Unmarshal(response, &networks)
	if err != nil {
		klog.Errorf("failed to unmarshal networks, err: %+v", err)
		return err
	}
	err = json.Unmarshal(response, &networks)
	if err != nil {
		klog.Errorf("failed to unmarshal networks, err: %+v", err)
		return err
	}
	for _, nw := range networks {
		if nw.Name == "azure" {
			proxy := hcn.L4ProxyPolicySetting{
				IP:          wr.MetadataIP,
				Port:        wr.MetadataPort,
				Protocol:    hcn.ProtocolTypeTCP,
				Destination: fmt.Sprintf("%s:%s", wr.RedirectIP, wr.RedirectPort),
				OutboundNAT: true,
			}
			setting, err := json.Marshal(proxy)
			if err != nil {
				klog.Errorf("failed to marshal proxy, err: %+v", err)
				return err
			}
			policy := hcn.NetworkPolicy{
				Type:     hcn.NetworkL4Proxy,
				Settings: setting,
			}
			networkRequest := hcn.PolicyNetworkRequest{
				Policies: []hcn.NetworkPolicy{policy},
			}
			policySetting, err := json.Marshal(networkRequest)
			if err != nil {
				klog.Errorf("failed to marshal policy setting, err: %+v", err)
				return err
			}
			requestSetting := hcn.ModifyNetworkSettingRequest{
				ResourceType: hcn.NetworkResourceTypePolicy,
				RequestType:  hcn.RequestTypeRemove,
				Settings:     policySetting,
			}
			klog.Infof("adding proxy policy to %s network id: %s", nw.Name, nw.Id)
			clientReq := msg.ModifyNetworkRequest{
				NetworkID:      nw.Id,
				RequestSetting: requestSetting,
			}
			clientR, err := json.Marshal(clientReq)
			if err != nil {
				klog.Errorf("failed to marshal client request, err: %+v", err)
				return err
			}
			clientRequest := msg.HNSRequest{
				Operation: msg.Modify,
				Entity:    msg.Network,
				Request:   clientR,
			}
			response, err = callHcnProxyAgentInternal(clientRequest)
			if err != nil {
				klog.Errorf("failed to modify, err: %+v", err)
				return err
			}
		}
	}
	return nil
}

func callHcnProxyAgentInternal(req msg.HNSRequest) ([]byte, error) {
	res := client.InvokeHNSRequest(req)
	if res.Error != nil {
		return nil, res.Error
	}

	b, _ := json.Marshal(res)
	klog.Infof("Server response: %s", string(b))

	return res.Response, nil
}
