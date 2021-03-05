package server

import (
	"github.com/Azure/aad-pod-identity/pkg/nmi/iptables"

	"k8s.io/klog/v2"
)

type LinuxRedirector struct {
	MetadataIP              string
	MetadataPort            string
	RedirectIP              string
	RedirectPort            string
	UpdateIntervalInSeconds int
}

func NewLinuxRedirector(metadataIP, metadatPort, redirectIP, redirectPort string, updateIntervalInSeconds int) *LinuxRedirector {
	return &LinuxRedirector{
		MetadataIP:              metadataIP,
		MetadataPort:            metadatPort,
		RedirectIP:              redirectIP,
		RedirectPort:            redirectPort,
		UpdateIntervalInSeconds: updateIntervalInSeconds,
	}
}

func (lr *LinuxRedirector) AddRedirectRules() error {
	return lr.updateIPTableRules()
}

func (lr *LinuxRedirector) RemoveRedirectRules() error {
	// clean up iptables rules
	return iptables.DeleteCustomChain()
}

// updateIPTableRules ensures the correct iptable rules are set
// such that metadata requests are received by nmi assigned port
// NOT originating from HostIP destined to metadata endpoint are
// routed to NMI endpoint
func (lr *LinuxRedirector) updateIPTableRules() error {
	klog.V(5).Infof("ip(%s) metadata address(%s:%s) nmi port(%s)", lr.RedirectIP, lr.MetadataIP, lr.MetadataPort, lr.RedirectPort)

	if err := iptables.AddCustomChain(lr.MetadataIP, lr.MetadataPort, lr.RedirectIP, lr.RedirectPort); err != nil {
		return err
	}
	if err := iptables.LogCustomChain(); err != nil {
		return err
	}
	return nil
}
