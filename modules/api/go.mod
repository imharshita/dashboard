module k8c.io/dashboard/v2

go 1.19

require (
	code.cloudfoundry.org/go-pubsub v0.0.0-20221028230532-6f1c373813c9
	github.com/Azure/azure-sdk-for-go v67.1.0+incompatible
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.2.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork v1.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions v1.0.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription v1.0.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Masterminds/semver/v3 v3.2.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.47
	github.com/aws/aws-sdk-go-v2 v1.17.1
	github.com/aws/aws-sdk-go-v2/config v1.18.3
	github.com/aws/aws-sdk-go-v2/credentials v1.13.3
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.73.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.23.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.23
	github.com/aws/aws-sdk-go-v2/service/sts v1.17.5
	github.com/aws/smithy-go v1.13.4
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/cristim/ec2-instances-info v0.0.0-20220804141719-da3da500d587
	github.com/digitalocean/godo v1.91.1
	github.com/embik/nutanix-client-go v0.1.0
	github.com/evanphx/json-patch v5.6.0+incompatible
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.1
	github.com/go-logr/zapr v1.2.3
	github.com/go-openapi/errors v0.20.3
	github.com/go-openapi/runtime v0.24.2
	github.com/go-openapi/strfmt v0.21.3
	github.com/go-openapi/swag v0.22.3
	github.com/go-openapi/validate v0.22.0
	github.com/go-swagger/go-swagger v0.30.3
	github.com/go-test/deep v1.0.8
	github.com/google/go-cmp v0.5.9
	github.com/gophercloud/gophercloud v1.0.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/websocket v1.5.0
	github.com/hetznercloud/hcloud-go v1.37.0
	github.com/kubermatic/machine-controller v1.55.1-0.20221208132905-d40a14dc25a8
	github.com/minio/minio-go/v7 v7.0.44
	github.com/onsi/ginkgo v1.16.5
	github.com/open-policy-agent/frameworks/constraint v0.0.0-20221109005544-7de84dff5081 // v0.5.1
	github.com/open-policy-agent/gatekeeper v0.0.0-20220504234711-ecf609290e2e // v3.8.1
	github.com/packethost/packngo v0.29.0
	github.com/prometheus/client_golang v1.14.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace
	github.com/stretchr/testify v1.8.1
	github.com/vmware/go-vcloud-director/v2 v2.17.0
	github.com/vmware/govmomi v0.29.0
	go.anx.io/go-anxcloud v0.4.6
	go.uber.org/zap v1.23.0
	golang.org/x/oauth2 v0.2.0
	google.golang.org/api v0.103.0
	gopkg.in/square/go-jose.v2 v2.6.0
	gopkg.in/yaml.v3 v3.0.1
	k8c.io/kubeone v1.5.3
	k8c.io/operating-system-manager v1.1.2-0.20221201183812-4f7c5a687353
	k8c.io/reconciler v0.3.1
	k8s.io/api v0.25.4
	k8s.io/apiextensions-apiserver v0.25.4
	k8s.io/apimachinery v0.25.4
	k8s.io/apiserver v0.25.4
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/code-generator v0.25.4
	k8s.io/klog/v2 v2.80.1
	k8s.io/kubectl v0.25.4
	k8s.io/metrics v0.25.4
	k8s.io/utils v0.0.0-20221128185143-99ec85e7a448
	kubevirt.io/api v0.58.0
	kubevirt.io/containerized-data-importer-api v1.55.1
	sigs.k8s.io/controller-runtime v0.13.1
	sigs.k8s.io/controller-tools v0.10.0
	sigs.k8s.io/yaml v1.3.0
)

replace (
	k8s.io/api => k8s.io/api v0.25.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.25.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.25.4
	k8s.io/client-go => k8s.io/client-go v0.25.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.25.4
	k8s.io/code-generator => k8s.io/code-generator v0.25.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.25.4
	k8s.io/kubectl => k8s.io/kubectl v0.25.4
	k8s.io/kubelet => k8s.io/kubelet v0.25.4
	k8s.io/metrics => k8s.io/metrics v0.25.4
)

replace github.com/ajeddeloh/go-json => github.com/coreos/go-json v0.0.0-20220810161552-7cce03887f34

// overrides required for OPA constraints
replace (
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp => go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v0.20.0
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v0.20.0
	go.opentelemetry.io/proto/otlp => go.opentelemetry.io/proto/otlp v0.7.0
)

require k8c.io/kubermatic/v2 v2.21.1-0.20221208091954-b39911917701

require (
	cloud.google.com/go/compute v1.12.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.1.1 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.28 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.21 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v0.7.0 // indirect
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/sprig/v3 v3.2.2 // indirect
	github.com/PaesslerAG/gval v1.2.1 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1 // indirect
	github.com/ajeddeloh/go-json v0.0.0-20220810161552-7cce03887f34 // indirect
	github.com/ajeddeloh/yaml v0.0.0-20170912190910-6b94386aeefd // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr v1.4.10 // indirect
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.8 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cert-manager/cert-manager v1.10.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/container-linux-config-transpiler v0.9.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/ignition v0.35.0 // indirect
	github.com/coreos/locksmith v0.6.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/distribution/distribution/v3 v3.0.0-20221127105615-b9f9073d1a60 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/gobuffalo/flect v0.3.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/cel-go v0.12.5 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.2.1-0.20210504230335-f78f29fc09ea // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/huandu/xstrings v1.3.3 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jessevdk/go-flags v1.5.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nutanix-cloud-native/prism-go-client v0.3.4 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/openshift/api v0.0.0-20211217221424-8779abfbd571 // indirect
	github.com/openshift/custom-resource-status v1.1.2 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/peterhellberg/link v1.2.0 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.12.0 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/toqueteos/webbrowser v1.2.0 // indirect
	github.com/vincent-petithory/dataurl v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.11.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go4.org v0.0.0-20201209231011-d4a079459e60 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/term v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/time v0.2.0 // indirect
	golang.org/x/tools v0.3.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	helm.sh/helm/v3 v3.10.2 // indirect
	k8s.io/autoscaler/vertical-pod-autoscaler v0.12.0 // indirect
	k8s.io/component-base v0.25.4 // indirect
	k8s.io/gengo v0.0.0-20220613173612-397b4ae3bce7 // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-aggregator v0.25.4 // indirect
	k8s.io/kube-openapi v0.0.0-20221123214604-86e75ddd809a // indirect
	k8s.io/kubelet v0.25.4 // indirect
	kubevirt.io/controller-lifecycle-operator-sdk/api v0.2.4 // indirect
	sigs.k8s.io/gateway-api v0.5.1 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)
