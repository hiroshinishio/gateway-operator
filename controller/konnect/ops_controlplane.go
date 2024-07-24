package konnect

import (
	"context"
	"errors"
	"fmt"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	operatorv1alpha1 "github.com/kong/gateway-operator/api/v1alpha1"
	k8sutils "github.com/kong/gateway-operator/pkg/utils/kubernetes"
)

func createControlPlane(
	ctx context.Context,
	sdk *sdkkonnectgo.SDK,
	logger logr.Logger,
	cp *operatorv1alpha1.KonnectControlPlane,
) error {
	setKonnectLabels(cp, &cp.Spec)

	resp, err := sdk.ControlPlanes.CreateControlPlane(ctx, cp.Spec.CreateControlPlaneRequest)
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
				cp.GetGeneration(),
			),
			cp,
		)
		return errHandled
	}

	cp.Status.SetKonnectID(*resp.ControlPlane.ID)
	k8sutils.SetCondition(
		k8sutils.NewConditionWithGeneration(
			KonnectEntityProgrammedConditionType,
			metav1.ConditionTrue,
			KonnectEntityProgrammedReason,
			"",
			cp.GetGeneration(),
		),
		cp,
	)

	return nil
}

func deleteControlPlane(
	ctx context.Context,
	sdk *sdkkonnectgo.SDK,
	logger logr.Logger,
	cp *operatorv1alpha1.KonnectControlPlane,
) error {
	id := cp.GetKonnectStatus().GetKonnectID()
	if id == "" {
		return fmt.Errorf("can't remove %T without a Konnect ID", cp)
	}

	resp, err := sdk.ControlPlanes.DeleteControlPlane(ctx, id)
	if errHandled := handleResp[operatorv1alpha1.KonnectControlPlane](err, resp, DeleteOp); errHandled != nil {
		var sdkError *sdkerrors.SDKError
		if errors.As(errHandled, &sdkError) && sdkError.StatusCode == 404 {
			logger.Info("entity not found in Konnect, skipping delete",
				"op", DeleteOp, "type", cp.GetTypeName(), "id", id,
			)
			return nil
		}
		return FailedKonnectOpError[operatorv1alpha1.KonnectControlPlane]{
			Op:  DeleteOp,
			Err: errHandled,
		}
	}

	return nil
}

func updateControlPlane(
	ctx context.Context,
	sdk *sdkkonnectgo.SDK,
	logger logr.Logger,
	cp *operatorv1alpha1.KonnectControlPlane,
) error {
	id := cp.GetKonnectStatus().GetKonnectID()
	if id == "" {
		return fmt.Errorf("can't update %T without a Konnect ID", cp)
	}

	setKonnectLabels(cp, &cp.Spec)
	req := components.UpdateControlPlaneRequest{
		Name:        sdkkonnectgo.String(cp.Spec.Name),
		Description: cp.Spec.Description,
		AuthType:    (*components.UpdateControlPlaneRequestAuthType)(cp.Spec.AuthType),
		ProxyUrls:   cp.Spec.ProxyUrls,
		Labels:      cp.Spec.Labels,
	}

	resp, err := sdk.ControlPlanes.UpdateControlPlane(ctx, id, req)
	var sdkError *sdkerrors.NotFoundError
	if errors.As(err, &sdkError) {
		logger.Info("entity not found in Konnect, trying to recreate",
			"type", cp.GetTypeName(), "id", id,
		)
		if err := createControlPlane(ctx, sdk, logger, cp); err != nil {
			return FailedKonnectOpError[operatorv1alpha1.KonnectControlPlane]{
				Op:  UpdateOp,
				Err: err,
			}
		}
		// Create succeeded, status is already set to just return.
		return nil
	}

	if errHandled := handleResp[operatorv1alpha1.KonnectControlPlane](err, resp, UpdateOp); errHandled != nil {
		k8sutils.SetCondition(
			k8sutils.NewConditionWithGeneration(
				KonnectEntityProgrammedConditionType,
				metav1.ConditionFalse,
				"FailedToUpdate",
				errHandled.Error(),
				cp.GetGeneration(),
			),
			cp,
		)
		return FailedKonnectOpError[operatorv1alpha1.KonnectControlPlane]{
			Op:  UpdateOp,
			Err: errHandled,
		}
	}

	cp.Status.SetKonnectID(*resp.ControlPlane.ID)
	k8sutils.SetCondition(
		k8sutils.NewConditionWithGeneration(
			KonnectEntityProgrammedConditionType,
			metav1.ConditionTrue,
			KonnectEntityProgrammedReason,
			"",
			cp.GetGeneration(),
		),
		cp,
	)

	return nil
}
