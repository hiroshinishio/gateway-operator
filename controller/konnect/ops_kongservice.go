package konnect

import (
	"context"
	"errors"
	"fmt"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectgocomp "github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgoops "github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
	"github.com/go-logr/logr"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	operatorv1alpha1 "github.com/kong/gateway-operator/api/v1alpha1"
	k8sutils "github.com/kong/gateway-operator/pkg/utils/kubernetes"
)

func createService(
	ctx context.Context,
	sdk *sdkkonnectgo.SDK,
	logger logr.Logger,
	cl client.Client,
	svc *configurationv1alpha1.KongService,
) error {
	resp, err := sdk.Services.CreateService(ctx, svc.Status.ControlPlaneID, sdkkonnectgocomp.CreateService{
		URL:            svc.Spec.KongServiceAPISpec.URL,
		ConnectTimeout: svc.Spec.KongServiceAPISpec.ConnectTimeout,
		Enabled:        svc.Spec.KongServiceAPISpec.Enabled,
		Host:           svc.Spec.KongServiceAPISpec.Host,
		Name:           svc.Spec.KongServiceAPISpec.Name,
		Path:           svc.Spec.KongServiceAPISpec.Path,
		Port:           svc.Spec.KongServiceAPISpec.Port,
		Protocol:       svc.Spec.KongServiceAPISpec.Protocol,
		ReadTimeout:    svc.Spec.KongServiceAPISpec.ReadTimeout,
		Retries:        svc.Spec.KongServiceAPISpec.Retries,
		Tags:           svc.Spec.KongServiceAPISpec.Tags,
		TLSVerify:      svc.Spec.KongServiceAPISpec.TLSVerify,
		TLSVerifyDepth: svc.Spec.KongServiceAPISpec.TLSVerifyDepth,
		WriteTimeout:   svc.Spec.KongServiceAPISpec.WriteTimeout,
	})

	// TODO: handle already exists
	// Can't adopt it as it will cause conflicts between the controller
	// that created that entity and already manages it, hm
	if errHandled := handleResp[operatorv1alpha1.KonnectControlPlane](err, resp, CreateOp); errHandled != nil {
		k8sutils.SetCondition(
			k8sutils.NewConditionWithGeneration(
				KonnectEntityProgrammedConditionType,
				metav1.ConditionFalse,
				"FailedToCreate",
				errHandled.Error(),
				svc.GetGeneration(),
			),
			&svc.Status,
		)
		return errHandled
	}

	svc.Status.SetKonnectID(resp.Service.ID)
	k8sutils.SetCondition(
		k8sutils.NewConditionWithGeneration(
			KonnectEntityProgrammedConditionType,
			metav1.ConditionTrue,
			KonnectEntityProgrammedReason,
			"",
			svc.GetGeneration(),
		),
		&svc.Status,
	)

	return nil
}

func updateService(
	ctx context.Context,
	sdk *sdkkonnectgo.SDK,
	logger logr.Logger,
	cl client.Client,
	svc *configurationv1alpha1.KongService,
) error {
	// TODO(pmalek) handle other types of CP ref
	nnCP := types.NamespacedName{
		Namespace: svc.Spec.ControlPlaneRef.KonnectNamespacedRef.Namespace,
		Name:      svc.Spec.ControlPlaneRef.KonnectNamespacedRef.Name,
	}
	if nnCP.Namespace == "" {
		nnCP.Namespace = svc.Namespace
	}
	var cp operatorv1alpha1.KonnectControlPlane
	if err := cl.Get(ctx, nnCP, &cp); err != nil {
		return fmt.Errorf("failed to get KonnectControlPlane %s: for Service %s: %w",
			nnCP, client.ObjectKeyFromObject(svc), err,
		)
	}

	resp, err := sdk.Services.UpsertService(ctx, sdkkonnectgoops.UpsertServiceRequest{
		ControlPlaneID: cp.Status.KonnectID,
		ServiceID:      svc.Status.KonnectID,
		CreateService: sdkkonnectgocomp.CreateService{
			URL:            svc.Spec.KongServiceAPISpec.URL,
			ConnectTimeout: svc.Spec.KongServiceAPISpec.ConnectTimeout,
			Enabled:        svc.Spec.KongServiceAPISpec.Enabled,
			Host:           svc.Spec.KongServiceAPISpec.Host,
			Name:           svc.Spec.KongServiceAPISpec.Name,
			Path:           svc.Spec.KongServiceAPISpec.Path,
			Port:           svc.Spec.KongServiceAPISpec.Port,
			Protocol:       svc.Spec.KongServiceAPISpec.Protocol,
			ReadTimeout:    svc.Spec.KongServiceAPISpec.ReadTimeout,
			Retries:        svc.Spec.KongServiceAPISpec.Retries,
			Tags:           svc.Spec.KongServiceAPISpec.Tags,
			TLSVerify:      svc.Spec.KongServiceAPISpec.TLSVerify,
			TLSVerifyDepth: svc.Spec.KongServiceAPISpec.TLSVerifyDepth,
			WriteTimeout:   svc.Spec.KongServiceAPISpec.WriteTimeout,
		},
	})

	// TODO: handle already exists
	// Can't adopt it as it will cause conflicts between the controller
	// that created that entity and already manages it, hm
	if errHandled := handleResp[configurationv1alpha1.KongService](err, resp, UpdateOp); errHandled != nil {
		k8sutils.SetCondition(
			k8sutils.NewConditionWithGeneration(
				KonnectEntityProgrammedConditionType,
				metav1.ConditionFalse,
				"FailedToCreate",
				errHandled.Error(),
				svc.GetGeneration(),
			),
			&svc.Status,
		)
		return errHandled
	}

	svc.Status.SetKonnectID(resp.Service.ID)
	svc.Status.ControlPlaneID = cp.Status.KonnectID
	k8sutils.SetCondition(
		k8sutils.NewConditionWithGeneration(
			KonnectEntityProgrammedConditionType,
			metav1.ConditionTrue,
			KonnectEntityProgrammedReason,
			"",
			svc.GetGeneration(),
		),
		&svc.Status,
	)

	return nil
}

func deleteService(
	ctx context.Context,
	sdk *sdkkonnectgo.SDK,
	logger logr.Logger,
	cl client.Client,
	svc *configurationv1alpha1.KongService,
) error {
	id := svc.GetStatus().GetKonnectID()
	if id == "" {
		return fmt.Errorf("can't remove %T without a Konnect ID", svc)
	}

	resp, err := sdk.Services.DeleteService(ctx, svc.Status.ControlPlaneID, id)
	if errHandled := handleResp[configurationv1alpha1.KongService](err, resp, DeleteOp); errHandled != nil {
		var sdkError *sdkerrors.SDKError
		if errors.As(errHandled, &sdkError) && sdkError.StatusCode == 404 {
			logger.Info("entity not found in Konnect, skipping delete",
				"op", DeleteOp, "type", svc.GetTypeName(), "id", id,
			)
			return nil
		}
		return FailedKonnectOpError[configurationv1alpha1.KongService]{
			Op:  DeleteOp,
			Err: errHandled,
		}
	}

	return nil
}
