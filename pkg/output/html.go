package output

import (
	"encoding/json"
	"github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	"github.com/kubesphere/kubeeye/pkg/constant"
	"github.com/kubesphere/kubeeye/pkg/utils"
	"io"
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

	open, err := os.Open(path.Join(constant.ResultPath, resultName))
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
		list := getOpaList(results.Spec.OpaResult.ResourceResults)
		resultCollection[constant.Opa] = list
	}
	if results.Spec.PrometheusResult != nil {
		prometheus := getPrometheus(results.Spec.PrometheusResult)
		resultCollection[constant.Prometheus] = prometheus
	}
	if results.Spec.NodeInfoResult != nil {
		fileChange := getFileChange(results.Spec.NodeInfoResult)
		resultCollection[constant.FileChange] = fileChange
		sysctl := getSysctl(results.Spec.NodeInfoResult)
		resultCollection[constant.Sysctl] = sysctl
		systemd := getSystemd(results.Spec.NodeInfoResult)
		resultCollection[constant.Systemd] = systemd
		filter := getFileFilter(results.Spec.NodeInfoResult)
		resultCollection[constant.FileFilter] = filter
		command := getCommand(results.Spec.NodeInfoResult)
		resultCollection[constant.CustomCommand] = command
	}

	if results.Spec.ComponentResult != nil {
		component := getComponent(results.Spec.ComponentResult)
		resultCollection[constant.Component] = component
	}

	var ruleNumber [][]interface{}
	for key, val := range results.Spec.InspectRuleTotal {
		var issues = len(resultCollection[key])
		if issues > 0 {
			issues -= 1
		}
		ruleNumber = append(ruleNumber, []interface{}{key, val, issues})
	}

	data := map[string]interface{}{"title": results.Annotations[constant.AnnotationStartTime], "overview": ruleNumber, "details": resultCollection}

	return nil, data
}

func getOpaList(result []v1alpha2.ResourceResult) (opaList []renderNode) {
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

func getPrometheus(pro []v1alpha2.PrometheusResult) []renderNode {
	var prometheus []renderNode
	for _, p := range pro {
		value := renderNode{}
		value.Children = append(value.Children, renderNode{Text: p.Result})
		prometheus = append(prometheus, value)
	}
	return prometheus
}

func getFileChange(infoResult map[string]v1alpha2.NodeInfoResult) []renderNode {
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
	for k, v := range infoResult {
		for _, item := range v.FileChangeResult {
			if item.Issues != nil && len(item.Issues) > 0 {
				val := renderNode{
					Children: []renderNode{
						{Text: item.Path},
						{Text: item.FileName},
						{Text: k},
						{Text: strings.Join(item.Issues, ",")},
						{Text: string(item.Level)},
					},
				}
				villeinage = append(villeinage, val)
			}

		}

	}
	return villeinage
}

func getFileFilter(fileResult map[string]v1alpha2.NodeInfoResult) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true, Children: []renderNode{
		{Text: "name"},
		{Text: "Path"},
		{Text: "nodeName"},
		{Text: "Issues"},
		{Text: "level"}},
	}
	villeinage = append(villeinage, header)
	for k, v := range fileResult {
		for _, result := range v.FileFilterResult {
			for _, issue := range result.Issues {
				content2 := []renderNode{{Text: result.FileName}, {Text: result.Path}, {Text: k}, {Text: issue}, {Text: string(result.Level)}}
				villeinage = append(villeinage, renderNode{Children: content2})
			}

		}

	}

	return villeinage
}
func getComponent(component []v1alpha2.ComponentResultItem) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true, Children: []renderNode{
		{Text: "name"},
		{Text: "namespace"},
		{Text: "endpoint"}},
	}
	villeinage = append(villeinage, header)

	for _, c := range component {
		value := []renderNode{{Text: c.Name}, {Text: c.Namespace}, {Text: c.Endpoint}}
		villeinage = append(villeinage, renderNode{Children: value})
	}

	return villeinage
}

func getSysctl(infoResult map[string]v1alpha2.NodeInfoResult) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "value"},
		}}
	villeinage = append(villeinage, header)
	for k, v := range infoResult {

		for _, item := range v.SysctlResult {
			if item.Assert {
				val := renderNode{
					Issues: item.Assert,
					Children: []renderNode{
						{Text: item.Name},
						{Text: k},
						{Text: *item.Value},
					}}
				villeinage = append(villeinage, val)
			}

		}

	}
	return villeinage
}

func getSystemd(infoResult map[string]v1alpha2.NodeInfoResult) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "value"},
		},
	}
	villeinage = append(villeinage, header)
	for k, v := range infoResult {
		for _, item := range v.SystemdResult {
			if item.Assert {
				val := renderNode{
					Issues: item.Assert,
					Children: []renderNode{
						{Text: item.Name},
						{Text: k},
						{Text: *item.Value},
					}}
				villeinage = append(villeinage, val)
			}
		}

	}
	return villeinage
}
func getCommand(commandResult map[string]v1alpha2.NodeInfoResult) []renderNode {
	var villeinage []renderNode
	header := renderNode{Header: true,
		Children: []renderNode{
			{Text: "name"},
			{Text: "nodeName"},
			{Text: "value"},
		},
	}
	villeinage = append(villeinage, header)
	for k, v := range commandResult {
		for _, item := range v.CommandResult {
			if item.Assert {
				val := renderNode{
					Issues: item.Assert,
					Children: []renderNode{
						{Text: item.Name},
						{Text: k},
						{Text: utils.BoolToString(item.Assert)},
					}}
				villeinage = append(villeinage, val)
			}
		}

	}
	return villeinage
}
