package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	"github.com/kubesphere/kubeeye/pkg/constant"
	"github.com/kubesphere/kubeeye/pkg/template"
	"github.com/kubesphere/kubeeye/pkg/utils"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"os"
	"path"
	"strings"
)

type renderNode struct {
	Text     string
	Issues   bool
	Header   bool
	Children []renderNode
}

func HtmlOut(resultName string) (error, map[string]interface{}) {

	var results v1alpha2.InspectResult

	open, err := os.Open(path.Join(constant.ResultPathPrefix, resultName))
	if err != nil {
		return err, nil
	}
	defer open.Close()

	all, err := io.ReadAll(open)
	if err != nil {
		return err, nil
	}

	err = json.Unmarshal(all, &results)
	if err != nil {
		return err, nil
	}
	var resultCollection = make(map[string][]renderNode, 5)

	if results.Spec.OpaResult.ResourceResults != nil {
		list := GetOpaList(results.Spec.OpaResult.ResourceResults)
		resultCollection[constant.Opa] = list
	}
	if results.Spec.PrometheusResult != nil {
		prometheus := GetPrometheus(results.Spec.PrometheusResult)
		resultCollection[constant.Prometheus] = prometheus
	}

	if results.Spec.FileChangeResult != nil {
		resultCollection[constant.FileChange] = GetFileChange(results.Spec.FileChangeResult)
	}

	if results.Spec.SysctlResult != nil {
		resultCollection[constant.Sysctl] = GetSysctl(results.Spec.SysctlResult)

	}
	if results.Spec.SystemdResult != nil {
		resultCollection[constant.Systemd] = GetSystemd(results.Spec.SystemdResult)

	}
	if results.Spec.FileFilterResult != nil {
		resultCollection[constant.FileFilter] = GetFileFilter(results.Spec.FileFilterResult)

	}

	if results.Spec.CommandResult != nil {
		resultCollection[constant.CustomCommand] = GetCommand(results.Spec.CommandResult)

	}
	if results.Spec.NodeInfo != nil {
		resultCollection[constant.NodeInfo] = GetNodeInfo(results.Spec.NodeInfo)
	}

	if results.Spec.ServiceConnectResult != nil {
		component := GetServiceConnect(results.Spec.ServiceConnectResult)
		resultCollection[constant.ServiceConnect] = component
	}

	var ruleNumber [][]interface{}
	for key, val := range results.Spec.InspectRuleTotal {
		var issues = len(resultCollection[key])
		if issues > 0 {
			issues -= 1
		}
		ruleNumber = append(ruleNumber, []interface{}{key, val, issues})
	}

	if os.Getenv("DISABLE_SYSTEM_COMPONENT") == "true" {
		delete(resultCollection, constant.Component)
	}

	data := map[string]interface{}{"title": results.Annotations[constant.AnnotationStartTime], "overview": ruleNumber, "details": resultCollection}

	if os.Getenv("DISABLE_OVERVIEW") == "true" {
		data = map[string]interface{}{"title": results.Annotations[constant.AnnotationStartTime], "details": resultCollection}
	}

	return nil, data
}

func GetOpaList(result []v1alpha2.ResourceResult) (opaList []renderNode) {
	opaList = append(opaList, renderNode{Header: true, Children: []renderNode{
		{Text: "Name"}, {Text: "Kind"}, {Text: "NameSpace"}, {Text: "Message"}, {Text: "Reason"}, {Text: "Level"},
	}})
	for _, resourceResult := range result {

		for _, item := range resourceResult.ResultItems {
			items := []renderNode{
				{Text: resourceResult.Name},
				{Text: resourceResult.ResourceType},
				{Text: resourceResult.NameSpace},
				{Text: item.Message},
				{Text: item.Reason},
				{Text: item.Level},
			}
			opaList = append(opaList, renderNode{Children: items})
		}
	}

	return opaList
}

func GetPrometheus(pro []v1alpha2.PrometheusResult) []renderNode {
	var prometheus []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "result"},
		}}

	prometheus = append(prometheus, header)

	for _, p := range pro {
		val := renderNode{
			Children: []renderNode{
				{Text: p.Name},
				{Text: p.Result},
			},
		}

		prometheus = append(prometheus, val)
	}
	return prometheus
}

func GetFileChange(fileChange []v1alpha2.FileChangeResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "path"},
			{Text: "nodeName"},
			{Text: "value"},
			{Text: "level"},
		}}
	villeinage = append(villeinage, header)

	for _, item := range fileChange {
		if item.Issues != nil && len(item.Issues) > 0 {
			val := renderNode{
				Children: []renderNode{
					{Text: item.Path},
					{Text: item.Name},
					{Text: item.NodeName},
					{Text: strings.Join(item.Issues, ",")},
					{Text: string(item.Level)},
				},
			}
			villeinage = append(villeinage, val)
		}

	}

	return villeinage
}

func GetFileFilter(fileResult []v1alpha2.FileChangeResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true, Children: []renderNode{
		{Text: "name"},
		{Text: "Path"},
		{Text: "nodeName"},
		{Text: "Issues"},
		{Text: "level"}},
	}
	villeinage = append(villeinage, header)

	for _, result := range fileResult {
		for _, issue := range result.Issues {
			content2 := []renderNode{{Text: result.Name}, {Text: result.Path}, {Text: result.NodeName}, {Text: issue}, {Text: string(result.Level)}}
			villeinage = append(villeinage, renderNode{Children: content2})
		}

	}

	return villeinage
}
func GetServiceConnect(component []v1alpha2.ServiceConnectResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true, Children: []renderNode{
		{Text: "name"},
		{Text: "namespace"},
		{Text: "endpoint"}},
	}
	villeinage = append(villeinage, header)

	for _, c := range component {
		if c.Assert {
			value := []renderNode{{Text: c.Name}, {Text: c.Namespace}, {Text: c.Endpoint}}
			villeinage = append(villeinage, renderNode{Children: value})
		}
	}

	return villeinage
}

func GetSysctl(sysctlResult []v1alpha2.NodeMetricsResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "value"},
		}}
	villeinage = append(villeinage, header)

	for _, item := range sysctlResult {
		if item.Assert {
			val := renderNode{
				Issues: item.Assert,
				Children: []renderNode{
					{Text: item.Name},
					{Text: item.NodeName},
					{Text: *item.Value},
				}}
			villeinage = append(villeinage, val)
		}
	}

	return villeinage
}

func GetNodeInfo(nodeInfo []v1alpha2.NodeInfoResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "resourcesType"},
			{Text: "mount"},
			{Text: "value"},
		}}
	villeinage = append(villeinage, header)

	for _, item := range nodeInfo {
		if item.Assert {
			val := renderNode{
				Issues: item.Assert,
				Children: []renderNode{
					{Text: item.Name},
					{Text: item.NodeName},
					{Text: item.ResourcesType.Type},
					{Text: item.ResourcesType.Mount},
					{Text: item.Value},
				}}
			villeinage = append(villeinage, val)
		}

	}

	return villeinage
}

func GetSystemd(systemdResult []v1alpha2.NodeMetricsResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "value"},
		},
	}
	villeinage = append(villeinage, header)

	for _, item := range systemdResult {
		if item.Assert {
			val := renderNode{
				Issues: item.Assert,
				Children: []renderNode{
					{Text: item.Name},
					{Text: item.NodeName},
					{Text: *item.Value},
				}}
			villeinage = append(villeinage, val)
		}
	}

	return villeinage
}
func GetCommand(commandResult []v1alpha2.CommandResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "value"},
		},
	}
	villeinage = append(villeinage, header)

	for _, item := range commandResult {
		if item.Assert {
			val := renderNode{
				Issues: item.Assert,
				Children: []renderNode{
					{Text: item.Name},
					{Text: item.NodeName},
					{Text: utils.BoolToString(item.Assert)},
				}}
			villeinage = append(villeinage, val)
		}
	}

	return villeinage
}

func GetNodesStatus(result *corev1.NodeList) (nodeList []renderNode) {
	nodeList = append(nodeList, renderNode{Header: true, Children: []renderNode{
		{Text: "Name"}, {Text: "Address"}, {Text: "Version"}, {Text: "Status"},
	}})
	for _, resourceResult := range result.Items {
		address := ""
		status := ""
		for _, addr := range resourceResult.Status.Addresses {
			if addr.Type == corev1.NodeInternalIP {
				address = addr.Address
			}
		}
		for _, condition := range resourceResult.Status.Conditions {
			if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
				status = "Ready"
			} else {
				status = "NotReady"
			}
		}
		items := []renderNode{
			{Text: resourceResult.Name},
			{Text: address},
			{Text: resourceResult.Status.NodeInfo.KubeletVersion},
			{Text: status},
		}

		nodeList = append(nodeList, renderNode{Children: items})
	}

	return nodeList
}

func GetAbnormalPods(result *corev1.PodList) (podList []renderNode) {
	podList = append(podList, renderNode{Header: true, Children: []renderNode{
		{Text: "Name"}, {Text: "Namespace"}, {Text: "Phase"}, {Text: "Reason"}, {Text: "Message"},
	}})
	for _, resourceResult := range result.Items {
		if resourceResult.Status.Phase != corev1.PodRunning && resourceResult.Status.Phase != corev1.PodSucceeded {
			reason := ""
			message := ""
			if len(resourceResult.Status.InitContainerStatuses) > 0 {
				for _, intContainer := range resourceResult.Status.InitContainerStatuses {
					if intContainer.State.Waiting != nil {
						if reason != "" {
							reason = reason + ", " + intContainer.State.Waiting.Reason
						} else {
							reason = intContainer.State.Waiting.Reason
						}
						if message != "" {
							message = message + "\n" + intContainer.State.Waiting.Message
						} else {
							message = intContainer.State.Waiting.Message
						}
					} else if intContainer.State.Terminated != nil {
						if reason != "" {
							reason = reason + ", " + intContainer.State.Terminated.Reason
						} else {
							reason = intContainer.State.Terminated.Reason
						}
						if message != "" {
							message = message + "\n" + intContainer.State.Terminated.Message
						} else {
							message = intContainer.State.Terminated.Message
						}
					}
				}
			}
			if len(resourceResult.Status.ContainerStatuses) > 0 {
				for _, container := range resourceResult.Status.ContainerStatuses {
					if container.State.Waiting != nil {
						if reason != "" {
							reason = reason + ", " + container.State.Waiting.Reason
						} else {
							reason = container.State.Waiting.Reason
						}
						if message != "" {
							message = message + "\n" + container.State.Waiting.Message
						} else {
							message = container.State.Waiting.Message
						}
					} else if container.State.Terminated != nil {
						if reason != "" {
							reason = reason + ", " + container.State.Terminated.Reason
						} else {
							reason = container.State.Terminated.Reason
						}
						if message != "" {
							message = message + "\n" + container.State.Terminated.Message
						} else {
							message = container.State.Terminated.Message
						}
					}
				}
			}
			items := []renderNode{
				{Text: resourceResult.Name},
				{Text: resourceResult.Namespace},
				{Text: string(resourceResult.Status.Phase)},
				{Text: reason},
				{Text: message},
			}
			podList = append(podList, renderNode{Children: items})
		}
	}

	return podList
}

func GenerateHtml(resultName string) error {
	htmlTemplate, err := template.GetInspectResultHtmlTemplate()
	if err != nil {
		return errors.Wrap(err, "GetInspectResultHtmlTemplate error")

	}
	err, m := HtmlOut(resultName)
	if err != nil {
		return errors.Wrap(err, "HtmlOut error")
	}
	data := bytes.NewBufferString("")
	err = htmlTemplate.Execute(data, m)
	if err != nil {
		return errors.Wrap(err, "render html template error")
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s.html", path.Join(constant.ResultPathPrefix, resultName)), data.Bytes(), 0644)
	if err != nil {
		return errors.Wrap(err, "write html file error")
	}
	return nil
}
