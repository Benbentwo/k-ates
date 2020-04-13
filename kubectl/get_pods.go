package kubectl

import (
	"github.com/Benbentwo/k-ates/templates"
	"github.com/Benbentwo/k-ates/util"
	"net/http"
)

//
// import (
// 	"context"
// 	"fmt"
// 	"github.com/Benbentwo/k-ates/util"
// 	v1 "k8s.io/api/core/v1"
// 	"k8s.io/apimachinery/pkg/api/errors"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// )
//
//
//
// func (kubectl *Kubectl)GetPods(namespace string) (*v1.PodList, error) {
// 	clientset := kubectl.Client
// 	// get pods in all the namespaces by omitting namespace
// 	// Or specify namespace to get pods in particular namespace
// 	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	if errors.IsNotFound(err) {
// 		fmt.Printf("Pod example-xxxxx not found in default namespace\n")
// 	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
// 		fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
// 	} else if err != nil {
// 		panic(err.Error())
// 	} else {
// 		fmt.Printf("Found example-xxxxx pod in default namespace\n")
// 	}
// 	util.Log(util.INFO, util.ColorInfo("There are "), util.ColorBold(len(pods.Items)), util.ColorInfo(" pods"))
// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
// 	// for {
// 	//
// 	// 	// Examples for error handling:
// 	// 	// - Use helper functions e.g. errors.IsNotFound()
// 	// 	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
// 	// 	_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
//
// 	//
// 	// 	time.Sleep(10 * time.Second)
// 	// }
// 	return pods, nil
// }

func GetPodsHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := util.GetRootPath()
	filepath := r.URL.Path[len(rootPath):]
	Log(DEBUG, util.ColorInfo("FilePath: \"")+filepath+util.ColorInfo("\""))
	Log(DEBUG, util.ColorInfo("RootPath: \"")+rootPath+util.ColorInfo("\""))

	Log(DEBUG, "Display Kubectl Get Pods Page")
	util.LoadTemplate(w, templates.K8, templates.K8Template{
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
			{
				Href: ".",
				Text: "Refresh",
			},
		},
		Headers: []string{
			// "A", "B",
		},
	})

	// t, _ := template.New("table").Parse(templates.Table)
	// t.Execute(w, table)
}
