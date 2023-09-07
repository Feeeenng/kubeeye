package inspect

import (
	"context"
	"encoding/json"
	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	"github.com/kubesphere/kubeeye/pkg/conf"
	"github.com/kubesphere/kubeeye/pkg/constant"
	"github.com/kubesphere/kubeeye/pkg/kube"
	"github.com/kubesphere/kubeeye/pkg/rules"
	"github.com/kubesphere/kubeeye/pkg/template"
	"github.com/prometheus/client_golang/api"
	apiprometheusv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"strings"
	"time"
)

type prometheusInspect struct {
}

func init() {
	RuleOperatorMap[constant.Prometheus] = &prometheusInspect{}
}

func (o *prometheusInspect) CreateJobTask(ctx context.Context, clients *kube.KubernetesClient, jobRule *kubeeyev1alpha2.JobRule, task *kubeeyev1alpha2.InspectTask, config *conf.JobConfig) (*kubeeyev1alpha2.JobPhase, error) {

	job := template.InspectJobsTemplate(config, jobRule.JobName, task, "", nil, constant.Prometheus)

	_, err := clients.ClientSet.BatchV1().Jobs(constant.DefaultNamespace).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		klog.Errorf("Failed to create Jobs  for node name:%s,err:%s", job.Name, err)
		return nil, err
	}
	return &kubeeyev1alpha2.JobPhase{JobName: jobRule.JobName, Phase: kubeeyev1alpha2.PhaseRunning}, err

}

func (o *prometheusInspect) RunInspect(ctx context.Context, ruless []kubeeyev1alpha2.JobRule, clients *kube.KubernetesClient, currentJobName string, ownerRef ...metav1.OwnerReference) ([]byte, error) {

	//_, exist, phase := utils.ArrayFinds(ruless, func(m kubeeyev1alpha2.JobRule) bool {
	//	return m.JobName == currentJobName
	//})
	//
	//if !exist {
	//	return nil, nil
	//}

	get, _ := clients.VersionClientSet.KubeeyeV1alpha2().InspectRules().Get(ctx, "prometheus-sample", metav1.GetOptions{})
	rule, _ := rules.MergeRule(*get)

	//var proRules []kubeeyev1alpha2.PrometheusRule
	//err := json.Unmarshal(phase.RunRule, &proRules)
	//if err != nil {
	//	klog.Error(err, " Failed to marshal kubeeye result")
	//	return nil, err
	//}

	var proRuleResult []kubeeyev1alpha2.PrometheusResult
	for _, proRule := range rule.Prometheus {
		proClient, err := api.NewClient(api.Config{
			Address: proRule.Endpoint,
		})
		if err != nil {
			klog.Error("create prometheus client failed", err)
			continue
		}
		queryApi := apiprometheusv1.NewAPI(proClient)
		query, _, err := queryApi.Query(ctx, *proRule.Rule, time.Now())
		if err != nil {
			klog.Errorf("failed to query rule:%s", *proRule.Rule)
			return nil, err
		}
		//fmt.Println(query.String())
		//marshal, err := json.Marshal(query)
		//if err != nil {
		//	klog.Error("marshal modal Samples failed", err)
		//	continue
		//}
		//var queryResults model.Samples
		//err = json.Unmarshal(marshal, &queryResults)
		//if err != nil {
		//	klog.Error("unmarshal modal Samples failed", err)
		//	continue
		//}
		//if queryResults.Len() == 0 {
		//	continue
		//}
		//var queryResultsMap []map[string]string
		//for _, result := range queryResults {
		//	temp := map[string]string{"value": result.Value.String(), "time": result.Timestamp.String(), "level": string(proRule.Level)}
		//	for name, value := range result.Metric {
		//		temp[formatName(name)] = string(value)
		//	}
		//	queryResultsMap = append(queryResultsMap, temp)
		//
		//}

		proRuleResult = append(proRuleResult, kubeeyev1alpha2.PrometheusResult{
			Result: query.String(),
			Level:  proRule.Level,
		})
	}

	marshal, err := json.Marshal(proRuleResult)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (o *prometheusInspect) GetResult(runNodeName string, resultCm *corev1.ConfigMap, resultCr *kubeeyev1alpha2.InspectResult) (*kubeeyev1alpha2.InspectResult, error) {
	var prometheus []kubeeyev1alpha2.PrometheusResult

	err := json.Unmarshal(resultCm.BinaryData[constant.Data], &prometheus)
	if err != nil {
		return nil, err
	}
	if prometheus == nil {
		return resultCr, nil
	}

	resultCr.Spec.PrometheusResult = prometheus

	return resultCr, nil
}

func formatName(name model.LabelName) string {
	return strings.Trim(string(name), "_")
}
