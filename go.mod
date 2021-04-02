module github.com/Azure/aad-pod-identity

go 1.16

require (
	contrib.go.opencensus.io/exporter/prometheus v0.2.0
	github.com/Azure/azure-sdk-for-go v52.3.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.18
	github.com/Azure/go-autorest/autorest/adal v0.9.13
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Microsoft/hcnproxy v1.0.2-0.20210104194639-98a553fdefaa
	github.com/Microsoft/hcsshim v0.8.15
	github.com/coreos/go-iptables v0.5.0
	github.com/fsnotify/fsnotify v1.4.9
	github.com/google/go-cmp v0.5.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	k8s.io/component-base v0.20.4
	k8s.io/klog/v2 v2.6.0
)

replace (
	github.com/Microsoft/hcnproxy/pkg => ./pkg/hcnproxy/vendor/github.com/Microsoft/hcnproxy/pkg
)