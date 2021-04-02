package types

import (
	"encoding/json"

	"github.com/Microsoft/hcsshim/hcn"
)

const (
	PipeName = `\\.\pipe\hnspipe`
	SSDL     = "D:P(A;;GA;;;BA)(A;;GA;;;SY)"
)

// OperationType specifies the type of HCN call to make.
type OperationType uint32

const (
	UnknownOperation OperationType = iota
	Create
	Delete
	Modify
	GetByName
	GetByID
	Enumerate
	GetHostName
)

// EntityType specifies the HCN entity on which to perform the operation.
type EntityType uint32

const (
	UnknownEntity EntityType = iota
	Network
	Endpoint
	EndpointV1
	LoadBalancer
	Host
)

// HNSRequest is the wrapper around a request from client
type HNSRequest struct {
	Operation OperationType
	Entity    EntityType
	Request   json.RawMessage
}

// HNSResponse is the wrapper around a response from client
type HNSResponse struct {
	Response json.RawMessage
	Error    error
}

// ModifyEndpointRequest stored the endpoint id and the setting to be used to modify endpoints
type ModifyEndpointRequest struct {
	EndpointID     string
	RequestSetting hcn.ModifyEndpointSettingRequest
}

// ModifyNetworkRequest stored the network id and the setting to be used to modify network
type ModifyNetworkRequest struct {
	NetworkID      string
	RequestSetting hcn.ModifyNetworkSettingRequest
}
