module github.com/Azure/aad-pod-identity/pkg/hcnproxy

go 1.16

replace github.com/microsoft/hcnproxy v0.0.0 => github.com/Microsoft/hcnproxy v1.0.2-0.20210104194639-98a553fdefaa // indirect

require (
	github.com/Microsoft/hcnproxy v1.0.1 // indirect
	github.com/microsoft/hcnproxy v0.0.0
)
