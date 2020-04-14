package kubectl

import (
	// "context"
	"fmt"
	"github.com/Benbentwo/k-ates/templates"
	"github.com/Benbentwo/k-ates/util"
	v1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func (kubectl *Kubectl) GetPods(namespace string) (*v1.PodList, error) {
	clientset := kubectl.Client
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		util.Log(util.ERROR, "Unable to fetch pods: ", err)
	}

	util.Log(util.INFO, util.ColorInfo("There are "), util.ColorBold(len(pods.Items)), util.ColorInfo(" pods"))
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	return pods, nil
}

func GetPodsHandler(w http.ResponseWriter, r *http.Request) {
	util.Log(util.INFO, "pod handler")
	rootPath := util.GetRootPath()
	kubectl := Kctl
	templateObject := templates.K8Template{
		Filename: "Pods",
		BreadCrumbs: []templates.ButtonLinks{
			{
				Text: "Kubectl",
				Href: rootPath + "kubectl",
			},
			{
				Text: "Get",
				Href: rootPath + "kubectl/get/",
			},
			{
				Text: "Pods",
				// Href: rootPath + "kubectl/get/pods/", // no Href means its current link
			},
		},
		Buttons: []templates.ButtonLinks{
			templates.RefreshButton,
			// {
			// 	Text: "Refresh Context",
			// 	Href: rootPath + "kubectl/context/",
			// },
		},
		Headers: []string{},
	}
	pods, err := kubectl.GetPods("")
	if err != nil {
		util.Log(util.ERROR, err)
		return
	}
	err = setDataToPods(&templateObject, pods)
	if err != nil {
		util.Log(util.ERROR, err)
		return
	}
	util.LoadTemplate(w, templates.K8, templateObject)

}

func setDataToPods(template *templates.K8Template, list *v1.PodList) error {
	headers, rows, err := PodListToTable(*list)
	if err != nil {
		return err
	}

	template.Rows = rows
	template.Headers = headers
	return nil
}
