//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuditResult) DeepCopyInto(out *AuditResult) {
	*out = *in
	in.Result.DeepCopyInto(&out.Result)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuditResult.
func (in *AuditResult) DeepCopy() *AuditResult {
	if in == nil {
		return nil
	}
	out := new(AuditResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraInfo) DeepCopyInto(out *ExtraInfo) {
	*out = *in
	if in.NamespacesList != nil {
		in, out := &in.NamespacesList, &out.NamespacesList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraInfo.
func (in *ExtraInfo) DeepCopy() *ExtraInfo {
	if in == nil {
		return nil
	}
	out := new(ExtraInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectPlan) DeepCopyInto(out *InspectPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectPlan.
func (in *InspectPlan) DeepCopy() *InspectPlan {
	if in == nil {
		return nil
	}
	out := new(InspectPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectPlanList) DeepCopyInto(out *InspectPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InspectPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectPlanList.
func (in *InspectPlanList) DeepCopy() *InspectPlanList {
	if in == nil {
		return nil
	}
	out := new(InspectPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectPlanSpec) DeepCopyInto(out *InspectPlanSpec) {
	*out = *in
	if in.Auditors != nil {
		in, out := &in.Auditors, &out.Auditors
		*out = make([]Auditor, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RuleNames != nil {
		in, out := &in.RuleNames, &out.RuleNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectPlanSpec.
func (in *InspectPlanSpec) DeepCopy() *InspectPlanSpec {
	if in == nil {
		return nil
	}
	out := new(InspectPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectPlanStatus) DeepCopyInto(out *InspectPlanStatus) {
	*out = *in
	in.LastScheduleTime.DeepCopyInto(&out.LastScheduleTime)
	in.NextScheduleTime.DeepCopyInto(&out.NextScheduleTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectPlanStatus.
func (in *InspectPlanStatus) DeepCopy() *InspectPlanStatus {
	if in == nil {
		return nil
	}
	out := new(InspectPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectRules) DeepCopyInto(out *InspectRules) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectRules.
func (in *InspectRules) DeepCopy() *InspectRules {
	if in == nil {
		return nil
	}
	out := new(InspectRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectRules) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectRulesList) DeepCopyInto(out *InspectRulesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InspectRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectRulesList.
func (in *InspectRulesList) DeepCopy() *InspectRulesList {
	if in == nil {
		return nil
	}
	out := new(InspectRulesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectRulesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectRulesSpec) DeepCopyInto(out *InspectRulesSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RuleItems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectRulesSpec.
func (in *InspectRulesSpec) DeepCopy() *InspectRulesSpec {
	if in == nil {
		return nil
	}
	out := new(InspectRulesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectRulesStatus) DeepCopyInto(out *InspectRulesStatus) {
	*out = *in
	in.ImportTime.DeepCopyInto(&out.ImportTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectRulesStatus.
func (in *InspectRulesStatus) DeepCopy() *InspectRulesStatus {
	if in == nil {
		return nil
	}
	out := new(InspectRulesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectTask) DeepCopyInto(out *InspectTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectTask.
func (in *InspectTask) DeepCopy() *InspectTask {
	if in == nil {
		return nil
	}
	out := new(InspectTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectTaskList) DeepCopyInto(out *InspectTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InspectTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectTaskList.
func (in *InspectTaskList) DeepCopy() *InspectTaskList {
	if in == nil {
		return nil
	}
	out := new(InspectTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InspectTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectTaskSpec) DeepCopyInto(out *InspectTaskSpec) {
	*out = *in
	if in.Auditors != nil {
		in, out := &in.Auditors, &out.Auditors
		*out = make([]Auditor, len(*in))
		copy(*out, *in)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]map[string]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectTaskSpec.
func (in *InspectTaskSpec) DeepCopy() *InspectTaskSpec {
	if in == nil {
		return nil
	}
	out := new(InspectTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InspectTaskStatus) DeepCopyInto(out *InspectTaskStatus) {
	*out = *in
	out.ClusterInfo = in.ClusterInfo
	if in.AuditResults != nil {
		in, out := &in.AuditResults, &out.AuditResults
		*out = make([]AuditResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.EndTimestamp != nil {
		in, out := &in.EndTimestamp, &out.EndTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InspectTaskStatus.
func (in *InspectTaskStatus) DeepCopy() *InspectTaskStatus {
	if in == nil {
		return nil
	}
	out := new(InspectTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeeyeAuditResult) DeepCopyInto(out *KubeeyeAuditResult) {
	*out = *in
	out.ScoreInfo = in.ScoreInfo
	if in.ResourceResults != nil {
		in, out := &in.ResourceResults, &out.ResourceResults
		*out = make([]ResourceResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ExtraInfo.DeepCopyInto(&out.ExtraInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeeyeAuditResult.
func (in *KubeeyeAuditResult) DeepCopy() *KubeeyeAuditResult {
	if in == nil {
		return nil
	}
	out := new(KubeeyeAuditResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceResult) DeepCopyInto(out *ResourceResult) {
	*out = *in
	if in.ResultItems != nil {
		in, out := &in.ResultItems, &out.ResultItems
		*out = make([]ResultItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceResult.
func (in *ResourceResult) DeepCopy() *ResourceResult {
	if in == nil {
		return nil
	}
	out := new(ResourceResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultItem) DeepCopyInto(out *ResultItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultItem.
func (in *ResultItem) DeepCopy() *ResultItem {
	if in == nil {
		return nil
	}
	out := new(ResultItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleItems) DeepCopyInto(out *RuleItems) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleItems.
func (in *RuleItems) DeepCopy() *RuleItems {
	if in == nil {
		return nil
	}
	out := new(RuleItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScoreInfo) DeepCopyInto(out *ScoreInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScoreInfo.
func (in *ScoreInfo) DeepCopy() *ScoreInfo {
	if in == nil {
		return nil
	}
	out := new(ScoreInfo)
	in.DeepCopyInto(out)
	return out
}
