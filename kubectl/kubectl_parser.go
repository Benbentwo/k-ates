package kubectl

import (
	v1 "k8s.io/api/core/v1"
	"strconv"
	"strings"
)

func PodListToTable(list v1.PodList) ([]string, [][]string, error) {
	data := make([][]string, len(list.Items))
	headers := getHeaders(list)
	for i, pod := range list.Items {
		data[i] = PodToStrArr(pod)
	}

	return headers, data, nil
}

func getHeaders(list v1.PodList) []string {
	headers := make([]string, 0)
	headers = append(headers, "Namespace")
	headers = append(headers, "Name")
	headers = append(headers, "Ready Containers")
	for _, pod := range list.Items {
		for _, container := range pod.Status.ContainerStatuses {
			if !container.Ready {
				headers = append(headers, "Not Ready")
			}
		}
	}
	headers = append(headers, "Status")
	headers = append(headers, "Restarts")
	headers = append(headers, "Labels")
	headers = append(headers, "Node")
	return headers

}
func PodToStrArr(pod v1.Pod) []string {

	data := make([]string, 0)
	data = append(data, pod.Namespace)

	data = append(data, pod.Name)

	notReady := make([]string, 0)
	ready := make([]string, 0)
	someNotReady := false
	for _, container := range pod.Status.ContainerStatuses {
		if container.Ready {
			ready = append(ready, container.Name)
		} else {
			someNotReady = true
			notReady = append(notReady, container.Name)
		}
	}

	data = append(data, strings.Join(ready, "\n"))

	if someNotReady {
		data = append(data, strings.Join(notReady, ","))
	}

	data = append(data, string(pod.Status.Phase))

	count := 0
	for _, container := range pod.Status.ContainerStatuses {
		count += int(container.RestartCount)
	}
	data = append(data, strconv.Itoa(count))

	labels := make([]string, 0)
	for key, value := range pod.ObjectMeta.Labels {
		labels = append(labels, strings.TrimSpace(key)+"="+strings.TrimSpace(value))
	}
	data = append(data, strings.Join(labels, "\n"))

	data = append(data, pod.Spec.NodeName)

	return data
}
