package kubectl

import (
	"github.com/Benbentwo/k-ates/templates"
	"github.com/Benbentwo/k-ates/util"
	"net/http"
)

// import (
// 	"github.com/Benbentwo/k-ates/parsers"
// 	"github.com/Benbentwo/k-ates/templates"
// 	"github.com/Benbentwo/k-ates/util"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/rest"
// 	"net/http"
// 	"os"
// )
//
// func GetContext() *rest.Config{
// 	// creates the in-cluster config
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return config
//
//
// }
//
// func GetClientSet(config *rest.Config) *kubernetes.Clientset {
// 	// creates the clientset
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return clientset
// }
//
// type Kubectl struct {
// 	Client		kubernetes.Clientset
// }
//
// func NewKubectl() *Kubectl {
// 	kubectl := &Kubectl{
// 		Client: *GetClientSet(GetContext()),
// 	}
// 	return kubectl
// }
//
var (
	Log   = util.Log
	DEBUG = util.DEBUG
	INFO  = util.INFO
	WARN  = util.WARN
	ERROR = util.ERROR
)

const (
	GetPodRoute = "kubectl/get/pods"
)

var rootPath = util.GetRootPath()
var TotalGetPodRoute = rootPath + GetPodRoute

func Handler(w http.ResponseWriter, r *http.Request) {
	util.LoadTemplate(w, templates.HOME, templates.Home{
		Filename: "Kubectl Home Page",
		Headers: []templates.ButtonLinks{
			{
				Text: "Get Pods",
				Href: TotalGetPodRoute,
			},
		},
	})

}
