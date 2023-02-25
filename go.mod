module github.com/Benbentwo/k-ates

go 1.13

require (
	github.com/TV4/logrus-stackdriver-formatter v0.1.0
	github.com/fatih/color v1.9.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang-collections/go-datastructures v0.0.0-20150211160725-59788d5eb259
	github.com/golang/protobuf v1.4.1 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rickar/props v0.0.0-20170718221555-0b06aeb2f037
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	k8s.io/api v0.0.0-20200429002227-34b661834646
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v0.0.0-20200429002950-b063729e49a6
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7
	k8s.io/api => k8s.io/api v0.0.0-20200429002227-34b661834646
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200429001640-17ad11ffb845
)
