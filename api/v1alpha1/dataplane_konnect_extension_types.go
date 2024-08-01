package v1alpha1

/*
Copyright 2024 Kong Inc.
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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&DataPlaneKonnectExtension{}, &DataPlaneKonnectExtensionList{})
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=kong;all
// +kubebuilder:subresource:status

// DataPlane is the Schema for the dataplanes API
type DataPlaneKonnectExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataPlaneKonnectExtensionSpec   `json:"spec,omitempty"`
	Status DataPlaneKonnectExtensionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataPlaneList contains a list of DataPlane
type DataPlaneKonnectExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataPlaneKonnectExtension `json:"items"`
}

// DataPlaneSpec defines the desired state of DataPlane
type DataPlaneKonnectExtensionSpec struct {
	// ControlPlaneID is the identifier of the Konnect Control Plane.
	// +kubebuilder:validation:Required
	ControlPlaneID string `json:"controlPlaneID"`

	// ControlPlaneRegion is the region of the Konnect Control Plane.
	// If not set, 'us' is used as the default region.
	// +optional
	// +kubebuilder:validation:Enum=us;eu
	// +kubebuilder:default=us
	ControlPlaneRegion *string `json:"controlPlaneRegion,omitempty"`

	// ClusterCertificate is a name of the Secret containing the Konnect Control Plane's cluster certificate.
	// +kubebuilder:validation:Required
	ClusterCertificate string `json:"clusterCertificate"`
}

// DataPlaneStatus defines the observed state of DataPlane
type DataPlaneKonnectExtensionStatus struct {
	// DataPlaneRefs is the array  of DataPlane references this is associated with.
	// A new reference is set by the operator when this extension is associated with
	// a DataPlane through its extensions spec.
	//
	// +kube:validation:Optional
	DataPlaneRefs []NamespacedRef `json:"dataPlaneRefs,omitempty"`
}
