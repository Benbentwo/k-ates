package kubectl

import (
	"flag"
	"github.com/Benbentwo/k-ates/templates"
	"github.com/Benbentwo/k-ates/util"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"path/filepath"
)

func GetContext() *rest.Config {
	// creates the in-cluster config
	util.Log(util.INFO, "Load In Cluster Config")
	config, err := rest.InClusterConfig()
	if err != nil {
		util.Log(util.ERROR, err)
		config = getClientSetFromOutOfCluster()
	}
	return config

}

// func LoadConfig() (*api.Config, *clientcmd.PathOptions, error) {
// 	po := clientcmd.NewDefaultPathOptions()
// 	if po == nil {
// 		return nil, po, fmt.Errorf("Could not find any default path options for the kubeconfig file usually found at ~/.kube/config")
// 	}
// 	config, err := po.GetStartingConfig()
// 	if err != nil {
// 		return nil, po, fmt.Errorf("Could not load the kube config file %s due to %s", po.GetDefaultFilename(), err)
// 	}
// 	return config, po, err
// }

func GetClientSet(config *rest.Config) *kubernetes.Clientset {
	// creates the clientset
	util.Log(util.DEBUG, "clientSetFunction")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		util.Log(util.ERROR, err)
		panic(err.Error())
	}
	return clientset
}
func getClientSetFromOutOfCluster() *rest.Config {
	util.Log(util.DEBUG, "Trying to load local config access")
	var kubeconfig *string
	if home := util.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig

	util.Log(util.DEBUG, *kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

type Kubectl struct {
	Client kubernetes.Clientset
}

var Kctl *Kubectl

func NewKubectl() *Kubectl {
	util.Log(util.DEBUG, "Creating new Kubectl")
	kubectl := &Kubectl{
		Client: *GetClientSet(GetContext()),
	}
	Kctl = kubectl
	return Kctl
}

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
	util.Log(util.INFO, "kubectl handler")
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

func HandleNewContext(_ http.ResponseWriter, _ *http.Request) {
	NewKubectl()
}
