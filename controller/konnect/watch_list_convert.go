package konnect

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	configurationv1 "github.com/kong/kubernetes-configuration/api/configuration/v1"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"

	operatorv1alpha1 "github.com/kong/gateway-operator/api/v1alpha1"
)

type SupportedObjectList interface {
	client.ObjectList
}

func ObjToObjList[T SupportedKonnectEntityType]() SupportedObjectList {
	var obj T
	switch any(obj).(type) {
	case operatorv1alpha1.KonnectControlPlane:
		return &operatorv1alpha1.KonnectControlPlaneList{}
	case configurationv1alpha1.KongService:
		return &configurationv1alpha1.KongServiceList{}
	case configurationv1alpha1.KongRoute:
		return &configurationv1alpha1.KongRouteList{}
	case configurationv1.KongConsumer:
		return &configurationv1.KongConsumerList{}
	default:
		panic(fmt.Sprintf("unsupported entity type %T", obj))
	}
}
