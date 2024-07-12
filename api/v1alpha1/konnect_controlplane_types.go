package v1alpha1

import (
	"context"

	sdkkonnectgocomp "github.com/Kong/sdk-konnect-go/models/components"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func init() {
	SchemeBuilder.Register(&KonnectControlPlane{}, &KonnectControlPlaneList{})
}

// KonnectControlPlane is the Schema for the konnectcontrolplanes API.
//
// +genclient
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed on Konnect",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
// +kubebuilder:printcolumn:name="ID",description="Konnect ID",type=string,JSONPath=`.status.id`
// +kubebuilder:printcolumn:name="OrgID",description="Konnect Organization ID this resource belongs to.",type=string,JSONPath=`.status.organizationID`
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.spec.konnectAPIAuthConfigurationRef) || has(oldSelf.spec.konnectAPIAuthConfigurationRef)", message="Konnect Configuration reference is immutable"
type KonnectControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KonnectControlPlaneSpec `json:"spec,omitempty"`

	Status configurationv1alpha1.KonnectEntityStatus `json:"status,omitempty"`
}

type KonnectControlPlaneSpec struct {
	sdkkonnectgocomp.CreateControlPlaneRequest `json:",inline"`

	KonnectConfiguration KonnectConfiguration `json:"konnect,omitempty"`
}

// GetKonnectStatus returns the Konnect Status of the KonnectControlPlane.
func (c *KonnectControlPlane) GetStatus() *configurationv1alpha1.KonnectEntityStatus {
	return &c.Status
}

func (c KonnectControlPlane) GetTypeName() string {
	return "KonnectControlPlane"
}

func (c *KonnectControlPlane) SetKonnectLabels(labels map[string]string) {
	c.Spec.Labels = labels
}

func (c *KonnectControlPlane) GetKonnectAPIAuthConfigurationRef() configurationv1alpha1.KonnectAPIAuthConfigurationRef {
	return c.Spec.KonnectConfiguration.APIAuthConfigurationRef
}

func (c *KonnectControlPlane) GetReconciliationWatchOptions(
	cl client.Client,
) []func(*ctrl.Builder) *ctrl.Builder {
	return []func(*ctrl.Builder) *ctrl.Builder{
		func(b *ctrl.Builder) *ctrl.Builder {
			// TODO(pmalek): this can be extracted and used in reconciler.go
			// as every Konnect entity will have a reference to the KonnectAPIAuthConfiguration.
			// This would require:
			// - mapping function from non List types to List types
			// - a function on each Konnect entity type to get the API Auth configuration
			//   reference from the object
			return b.Watches(
				&KonnectAPIAuthConfiguration{},
				handler.EnqueueRequestsFromMapFunc(
					enqueueKonnectControlPlaneForKonnectAPIAuthConfiguration(cl),
				),
			)
		},
	}
}

func enqueueKonnectControlPlaneForKonnectAPIAuthConfiguration(
	cl client.Client,
) func(ctx context.Context, obj client.Object) []reconcile.Request {
	return func(ctx context.Context, obj client.Object) []reconcile.Request {
		auth, ok := obj.(*KonnectAPIAuthConfiguration)
		if !ok {
			return nil
		}
		var l KonnectControlPlaneList
		if err := cl.List(ctx, &l); err != nil {
			return nil
		}
		var ret []reconcile.Request
		for _, cp := range l.Items {
			authRef := cp.GetKonnectAPIAuthConfigurationRef()
			if authRef.Name != auth.Name ||
				authRef.Namespace != auth.Namespace {
				continue
			}
			ret = append(ret, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: cp.Namespace,
					Name:      cp.Name,
				},
			})
		}
		return ret
	}
}

// +kubebuilder:object:root=true
type KonnectControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KonnectControlPlane `json:"items"`
}
