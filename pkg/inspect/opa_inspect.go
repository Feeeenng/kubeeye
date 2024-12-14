package inspect

import (
	"context"
	"encoding/json"
	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	"github.com/kubesphere/kubeeye/pkg/collectors/opa"
	"github.com/kubesphere/kubeeye/pkg/constant"
	"github.com/kubesphere/kubeeye/pkg/kube"
	"github.com/kubesphere/kubeeye/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/klog/v2"
)

type OpaInspect struct {
}

func init() {
	RuleOperatorMap[constant.Opa] = &OpaInspect{}
}

func (o *OpaInspect) RunInspect(ctx context.Context, rules []kubeeyev1alpha2.JobRule, clients *kube.KubernetesClient, currentJobName string, informers informers.SharedInformerFactory, ownerRef ...metav1.OwnerReference) ([]byte, error) {

	klog.Info("Fetching Rego Rules")

	_, exist, phase := utils.ArrayFinds(rules, func(m kubeeyev1alpha2.JobRule) bool {
		return m.JobName == currentJobName
	})

	if exist {
		rulesManager := opa.NewRulesManager()

		var opaRules []kubeeyev1alpha2.OpaRule
		err := json.Unmarshal(phase.RunRule, &opaRules)
		if err != nil {
			klog.Errorf("unmarshal opaRule failed,err:%s\n", err)
			return nil, err
		}

		klog.Info("Adding Rego Rules")

		for i := range opaRules {
			err := rulesManager.AddRule(&opaRules[i])
			if err != nil {
				klog.Errorf("add rule failed,err:%s\n", err)
				return nil, err
			}
		}

		klog.Info("Fetching resources")

		resourcesCollector, err := opa.NewResourceCollector(clients.KubeConfig)
		if err != nil {
			klog.Errorf("new resourceCollector failed,err:%s\n", err)
			return nil, err
		}

		resourcesManager := opa.NewResourcesManager()

		for key := range rulesManager.Rules {
			err := resourcesManager.AddResource(key, resourcesCollector)
			if err != nil {
				klog.Errorf("add resource failed,err:%s\n", err)
				return nil, err
			}
		}

		klog.Info("Checking Rego Rules")

		opaChecker := opa.NewOPAChecker(1000, 100)

		result, err := opaChecker.VailOpaRulesResult(rulesManager, resourcesManager)
		marshal, err := json.Marshal(result)

		if err != nil {
			klog.Error("marshal opaRule failed,err:%s\n", err)
			return nil, err
		}

		return marshal, nil
	}
	return nil, nil
}

func (o *OpaInspect) GetResult(runNodeName string, resultCm *corev1.ConfigMap, resultCr *kubeeyev1alpha2.InspectResult) (*kubeeyev1alpha2.InspectResult, error) {
	var opaResult kubeeyev1alpha2.KubeeyeOpaResult
	err := json.Unmarshal(resultCm.BinaryData[constant.Data], &opaResult)
	if err != nil {
		return nil, err
	}

	resultCr.Spec.OpaResult = opaResult

	return resultCr, nil
}
