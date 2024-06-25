package konnect

// TODO(pmalek): move this to API so that it's part of the API contract.
// TODO(pmalek): document this.

const (
	KonnectEntityProgrammedConditionType = "Programmed"

	KonnectEntityProgrammedReason = "Programmed"
)

const (
	KonnectAPIAuthConfigurationValidConditionType = "Valid"

	KonnectAPIAuthConfigurationReasonValid   = "Valid"
	KonnectAPIAuthConfigurationReasonInvalid = "Invalid"
)

const (
	ControlPlaneRefValidConditionType = "ControlPlaneRefValid"

	ControlPlaneRefReasonValid   = "Valid"
	ControlPlaneRefReasonInvalid = "Invalid"
)

const (
	KonnectEntityAPIAuthConfigurationRefValidConditionType = "APIAuthRefValid"

	KonnectEntityAPIAuthConfigurationRefReasonValid   = "Valid"
	KonnectEntityAPIAuthConfigurationRefReasonInvalid = "Invalid"
)
