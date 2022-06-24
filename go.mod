module github.com/Netflix/titus-executor

go 1.13

require (
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/DataDog/datadog-go v3.4.0+incompatible // indirect
	github.com/Datadog/opencensus-go-exporter-datadog v0.0.0-20190503082300-0f32ad59ab08
	github.com/Microsoft/go-winio v0.4.15 // indirect
	github.com/Netflix/metrics-client-go v0.0.0-20171019173821-bb173f41fc07
	github.com/Netflix/spectator-go v0.0.0-20190913215732-d4e0463555ef
	github.com/Netflix/titus-api-definitions v0.0.3-rc.21
	github.com/Netflix/titus-kube-common v0.28.0
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/alessio/shellescape v0.0.0-20190409004728-b115ca0f9053 // indirect
	github.com/apparentlymart/go-cidr v1.0.0
	github.com/aws/aws-sdk-go v1.40.6
	github.com/containernetworking/cni v0.8.1
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/coreos/go-systemd/v22 v22.3.2
	github.com/cyphar/filepath-securejoin v0.2.3
	github.com/docker/distribution v2.8.0+incompatible
	github.com/docker/docker v20.10.2+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0
	github.com/ftrvxmtrx/fd v0.0.0-20150925145434-c6d800382fff
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.5
	github.com/google/go-jsonnet v0.17.0
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/renameio v1.0.1
	github.com/google/uuid v1.1.2
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1
	github.com/jteeuwen/go-bindata v3.0.7+incompatible
	github.com/karlseguin/ccache/v2 v2.0.8
	github.com/leanovate/gopter v0.0.0-20170420174722-9e6101e5a875
	github.com/lib/pq v1.10.0
	github.com/m7shapan/cidr v0.0.0-20200427124835-7eba0889a5d2
	github.com/moby/sys/mountinfo v0.5.0
	github.com/mvisonneau/go-ebsnvme v0.0.0-20201026165225-e63797fabc2f
	github.com/myitcv/gobin v0.0.14
	github.com/netflix-skunkworks/opencensus-go-exporter-datadog v0.0.0-20190911150647-ef71dde58796
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v1.1.2
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/assertions v1.1.0 // indirect
	github.com/soheilhy/cmux v0.1.5
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli v1.22.5
	github.com/vishvananda/netlink v1.1.1-0.20210330154013-f5de75959ad5
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae
	github.com/wercker/journalhook v0.0.0-20180428041537-5d0a5ae867b3
	go.opencensus.io v0.22.3
	go.uber.org/multierr v1.6.0
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20211116061358-0a5406a5449c
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	golang.org/x/tools v0.1.2
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/alessio/shellescape.v1 v1.0.0-20170105083845-52074bc9df61
	gopkg.in/ini.v1 v1.52.0 // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.22.10
	k8s.io/apimachinery v0.22.10
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/utils v0.0.0-20211116205334-6203023598ed
)

replace github.com/docker/docker => github.com/moby/moby v0.0.0-20190408150954-50ebe4562dfc

replace sourcegraph.com/sqs/pbtypes => github.com/sqs/pbtypes v1.0.0

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.22.10

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.22.10

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.22.10

replace k8s.io/apiserver => k8s.io/apiserver v0.22.10

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.22.10

replace k8s.io/cri-api => k8s.io/cri-api v0.22.11-rc.0

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.22.10

replace k8s.io/kubelet => k8s.io/kubelet v0.22.10

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.22.10

replace k8s.io/apimachinery => k8s.io/apimachinery v0.22.11-rc.0

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.22.10

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.22.10

replace k8s.io/component-base => k8s.io/component-base v0.22.10

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.22.10

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.22.10

replace k8s.io/metrics => k8s.io/metrics v0.22.10

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.22.10

replace k8s.io/code-generator => k8s.io/code-generator v0.22.11-rc.0

replace k8s.io/client-go => k8s.io/client-go v0.22.10

replace k8s.io/kubectl => k8s.io/kubectl v0.22.10

replace k8s.io/api => k8s.io/api v0.22.10

replace k8s.io/component-helpers => k8s.io/component-helpers v0.22.10

replace k8s.io/controller-manager => k8s.io/controller-manager v0.22.10

replace k8s.io/mount-utils => k8s.io/mount-utils v0.22.11-rc.0

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.22.10

replace k8s.io/sample-controller => k8s.io/sample-controller v0.22.10

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.22.10
