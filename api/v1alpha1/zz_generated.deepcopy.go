//go:build !ignore_autogenerated

/*
Copyright 2022 Kong Inc.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AICloudProvider) DeepCopyInto(out *AICloudProvider) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AICloudProvider.
func (in *AICloudProvider) DeepCopy() *AICloudProvider {
	if in == nil {
		return nil
	}
	out := new(AICloudProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AICloudProviderAPITokenRef) DeepCopyInto(out *AICloudProviderAPITokenRef) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AICloudProviderAPITokenRef.
func (in *AICloudProviderAPITokenRef) DeepCopy() *AICloudProviderAPITokenRef {
	if in == nil {
		return nil
	}
	out := new(AICloudProviderAPITokenRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIGateway) DeepCopyInto(out *AIGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIGateway.
func (in *AIGateway) DeepCopy() *AIGateway {
	if in == nil {
		return nil
	}
	out := new(AIGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIGatewayConsumerRef) DeepCopyInto(out *AIGatewayConsumerRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIGatewayConsumerRef.
func (in *AIGatewayConsumerRef) DeepCopy() *AIGatewayConsumerRef {
	if in == nil {
		return nil
	}
	out := new(AIGatewayConsumerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIGatewayEndpoint) DeepCopyInto(out *AIGatewayEndpoint) {
	*out = *in
	if in.AvailableModels != nil {
		in, out := &in.AvailableModels, &out.AvailableModels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Consumer = in.Consumer
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIGatewayEndpoint.
func (in *AIGatewayEndpoint) DeepCopy() *AIGatewayEndpoint {
	if in == nil {
		return nil
	}
	out := new(AIGatewayEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIGatewayList) DeepCopyInto(out *AIGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AIGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIGatewayList.
func (in *AIGatewayList) DeepCopy() *AIGatewayList {
	if in == nil {
		return nil
	}
	out := new(AIGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AIGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIGatewaySpec) DeepCopyInto(out *AIGatewaySpec) {
	*out = *in
	if in.LargeLanguageModels != nil {
		in, out := &in.LargeLanguageModels, &out.LargeLanguageModels
		*out = new(LargeLanguageModels)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudProviderCredentials != nil {
		in, out := &in.CloudProviderCredentials, &out.CloudProviderCredentials
		*out = new(AICloudProviderAPITokenRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIGatewaySpec.
func (in *AIGatewaySpec) DeepCopy() *AIGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(AIGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AIGatewayStatus) DeepCopyInto(out *AIGatewayStatus) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]AIGatewayEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AIGatewayStatus.
func (in *AIGatewayStatus) DeepCopy() *AIGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(AIGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudHostedLargeLanguageModel) DeepCopyInto(out *CloudHostedLargeLanguageModel) {
	*out = *in
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		*out = new(string)
		**out = **in
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(LLMPromptType)
		**out = **in
	}
	if in.DefaultPrompts != nil {
		in, out := &in.DefaultPrompts, &out.DefaultPrompts
		*out = make([]LLMPrompt, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultPromptParams != nil {
		in, out := &in.DefaultPromptParams, &out.DefaultPromptParams
		*out = new(LLMPromptParams)
		(*in).DeepCopyInto(*out)
	}
	out.AICloudProvider = in.AICloudProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudHostedLargeLanguageModel.
func (in *CloudHostedLargeLanguageModel) DeepCopy() *CloudHostedLargeLanguageModel {
	if in == nil {
		return nil
	}
	out := new(CloudHostedLargeLanguageModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneMetricsExtension) DeepCopyInto(out *DataPlaneMetricsExtension) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneMetricsExtension.
func (in *DataPlaneMetricsExtension) DeepCopy() *DataPlaneMetricsExtension {
	if in == nil {
		return nil
	}
	out := new(DataPlaneMetricsExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataPlaneMetricsExtension) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneMetricsExtensionList) DeepCopyInto(out *DataPlaneMetricsExtensionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataPlaneMetricsExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneMetricsExtensionList.
func (in *DataPlaneMetricsExtensionList) DeepCopy() *DataPlaneMetricsExtensionList {
	if in == nil {
		return nil
	}
	out := new(DataPlaneMetricsExtensionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataPlaneMetricsExtensionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneMetricsExtensionSpec) DeepCopyInto(out *DataPlaneMetricsExtensionSpec) {
	*out = *in
	in.ServiceSelector.DeepCopyInto(&out.ServiceSelector)
	out.Config = in.Config
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneMetricsExtensionSpec.
func (in *DataPlaneMetricsExtensionSpec) DeepCopy() *DataPlaneMetricsExtensionSpec {
	if in == nil {
		return nil
	}
	out := new(DataPlaneMetricsExtensionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPlaneMetricsExtensionStatus) DeepCopyInto(out *DataPlaneMetricsExtensionStatus) {
	*out = *in
	if in.ControlPlaneRef != nil {
		in, out := &in.ControlPlaneRef, &out.ControlPlaneRef
		*out = new(NamespacedRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPlaneMetricsExtensionStatus.
func (in *DataPlaneMetricsExtensionStatus) DeepCopy() *DataPlaneMetricsExtensionStatus {
	if in == nil {
		return nil
	}
	out := new(DataPlaneMetricsExtensionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtensionRef) DeepCopyInto(out *ExtensionRef) {
	*out = *in
	in.NamespacedRef.DeepCopyInto(&out.NamespacedRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtensionRef.
func (in *ExtensionRef) DeepCopy() *ExtensionRef {
	if in == nil {
		return nil
	}
	out := new(ExtensionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfiguration) DeepCopyInto(out *KonnectAPIAuthConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfiguration.
func (in *KonnectAPIAuthConfiguration) DeepCopy() *KonnectAPIAuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectAPIAuthConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationList) DeepCopyInto(out *KonnectAPIAuthConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectAPIAuthConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationList.
func (in *KonnectAPIAuthConfigurationList) DeepCopy() *KonnectAPIAuthConfigurationList {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectAPIAuthConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationRef) DeepCopyInto(out *KonnectAPIAuthConfigurationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationRef.
func (in *KonnectAPIAuthConfigurationRef) DeepCopy() *KonnectAPIAuthConfigurationRef {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationSpec) DeepCopyInto(out *KonnectAPIAuthConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationSpec.
func (in *KonnectAPIAuthConfigurationSpec) DeepCopy() *KonnectAPIAuthConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectAPIAuthConfigurationStatus) DeepCopyInto(out *KonnectAPIAuthConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectAPIAuthConfigurationStatus.
func (in *KonnectAPIAuthConfigurationStatus) DeepCopy() *KonnectAPIAuthConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(KonnectAPIAuthConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectControlPlane) DeepCopyInto(out *KonnectControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectControlPlane.
func (in *KonnectControlPlane) DeepCopy() *KonnectControlPlane {
	if in == nil {
		return nil
	}
	out := new(KonnectControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectControlPlaneList) DeepCopyInto(out *KonnectControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KonnectControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectControlPlaneList.
func (in *KonnectControlPlaneList) DeepCopy() *KonnectControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(KonnectControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KonnectControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectControlPlaneSpec) DeepCopyInto(out *KonnectControlPlaneSpec) {
	*out = *in
	in.CreateControlPlaneRequest.DeepCopyInto(&out.CreateControlPlaneRequest)
	out.KonnectConfiguration = in.KonnectConfiguration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectControlPlaneSpec.
func (in *KonnectControlPlaneSpec) DeepCopy() *KonnectControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(KonnectControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMPrompt) DeepCopyInto(out *LLMPrompt) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(LLMPromptRole)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMPrompt.
func (in *LLMPrompt) DeepCopy() *LLMPrompt {
	if in == nil {
		return nil
	}
	out := new(LLMPrompt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMPromptParams) DeepCopyInto(out *LLMPromptParams) {
	*out = *in
	if in.Temperature != nil {
		in, out := &in.Temperature, &out.Temperature
		*out = new(string)
		**out = **in
	}
	if in.MaxTokens != nil {
		in, out := &in.MaxTokens, &out.MaxTokens
		*out = new(int)
		**out = **in
	}
	if in.TopK != nil {
		in, out := &in.TopK, &out.TopK
		*out = new(int)
		**out = **in
	}
	if in.TopP != nil {
		in, out := &in.TopP, &out.TopP
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMPromptParams.
func (in *LLMPromptParams) DeepCopy() *LLMPromptParams {
	if in == nil {
		return nil
	}
	out := new(LLMPromptParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LargeLanguageModels) DeepCopyInto(out *LargeLanguageModels) {
	*out = *in
	if in.CloudHosted != nil {
		in, out := &in.CloudHosted, &out.CloudHosted
		*out = make([]CloudHostedLargeLanguageModel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LargeLanguageModels.
func (in *LargeLanguageModels) DeepCopy() *LargeLanguageModels {
	if in == nil {
		return nil
	}
	out := new(LargeLanguageModels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfig) DeepCopyInto(out *MetricsConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfig.
func (in *MetricsConfig) DeepCopy() *MetricsConfig {
	if in == nil {
		return nil
	}
	out := new(MetricsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedRef) DeepCopyInto(out *NamespacedRef) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedRef.
func (in *NamespacedRef) DeepCopy() *NamespacedRef {
	if in == nil {
		return nil
	}
	out := new(NamespacedRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSelector) DeepCopyInto(out *ServiceSelector) {
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make([]ServiceSelectorEntry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector.
func (in *ServiceSelector) DeepCopy() *ServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSelectorEntry) DeepCopyInto(out *ServiceSelectorEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelectorEntry.
func (in *ServiceSelectorEntry) DeepCopy() *ServiceSelectorEntry {
	if in == nil {
		return nil
	}
	out := new(ServiceSelectorEntry)
	in.DeepCopyInto(out)
	return out
}
