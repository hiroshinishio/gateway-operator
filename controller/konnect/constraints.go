package konnect

import (
	"fmt"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	operatorv1alpha1.KonnectControlPlane | configurationv1alpha1.KongService | configurationv1alpha1.KongRoute
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

	GetStatus() *configurationv1alpha1.KonnectEntityStatus
	// GetStatusID() string
	// SetStatusID(string)
	// GetServerURL() string
	// SetServerURL(string)
	SetKonnectLabels(labels map[string]string)
	GetReconciliationWatchOptions(client.Client) []func(*ctrl.Builder) *ctrl.Builder
	GetKonnectAPIAuthConfigurationRef() configurationv1alpha1.KonnectAPIAuthConfigurationRef
}
