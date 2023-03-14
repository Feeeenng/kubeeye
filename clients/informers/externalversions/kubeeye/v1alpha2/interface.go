/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/kubesphere/kubeeye/clients/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// InspectPlans returns a InspectPlanInformer.
	InspectPlans() InspectPlanInformer
	// InspectRules returns a InspectRulesInformer.
	InspectRules() InspectRulesInformer
	// InspectTasks returns a InspectTaskInformer.
	InspectTasks() InspectTaskInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// InspectPlans returns a InspectPlanInformer.
func (v *version) InspectPlans() InspectPlanInformer {
	return &inspectPlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InspectRules returns a InspectRulesInformer.
func (v *version) InspectRules() InspectRulesInformer {
	return &inspectRulesInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InspectTasks returns a InspectTaskInformer.
func (v *version) InspectTasks() InspectTaskInformer {
	return &inspectTaskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
