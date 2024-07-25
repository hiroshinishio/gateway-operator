package konnect

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configurationv1 "github.com/kong/kubernetes-configuration/api/configuration/v1"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"

	operatorv1alpha1 "github.com/kong/gateway-operator/api/v1alpha1"
)

// TODO(pmalek): this could be useful to have a generic way to handle controller-runtime
// handlers/predicates but "sigs.k8s.io/controller-runtime/pkg/client".List is not generic
// and doesn't provide the actual list.
func ListTypeForType[T SupportedKonnectEntityType](e *T) client.ObjectList {
	switch any(e).(type) {
	case *operatorv1alpha1.KonnectControlPlane:
		return &operatorv1alpha1.KonnectControlPlaneList{}
	default:
		panic(fmt.Sprintf("unsupported entity type %T", e))
	}
}

type SupportedKonnectEntityType interface {
	operatorv1alpha1.KonnectControlPlane |
		configurationv1alpha1.KongService |
		configurationv1alpha1.KongRoute |
		configurationv1.KongConsumer |
		configurationv1alpha1.KongPluginBinding
	// TODO: add other types

	GetTypeName() string
}

type EntityType[
	T SupportedKonnectEntityType,
] interface {
	*T

	// Kubernetes Object methods

	GetObjectMeta() metav1.Object
	client.Object

	// Added methods

	GetConditions() []metav1.Condition
	SetConditions([]metav1.Condition)
	GetKonnectStatus() *configurationv1alpha1.KonnectEntityStatus
	// GetStatusID() string
	// SetStatusID(string)
	// GetServerURL() string
	// SetServerURL(string)

	// TODO(pmalek): not all entities can have labels.
	// SetKonnectLabels(labels map[string]string)

	GetKonnectAPIAuthConfigurationRef() configurationv1alpha1.KonnectAPIAuthConfigurationRef
}
