// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsDefenderApplicationControlSupplementalPolicyStatuses undocumented
type WindowsDefenderApplicationControlSupplementalPolicyStatuses int

const (
	// WindowsDefenderApplicationControlSupplementalPolicyStatusesVUnknown undocumented
	WindowsDefenderApplicationControlSupplementalPolicyStatusesVUnknown WindowsDefenderApplicationControlSupplementalPolicyStatuses = 0
	// WindowsDefenderApplicationControlSupplementalPolicyStatusesVSuccess undocumented
	WindowsDefenderApplicationControlSupplementalPolicyStatusesVSuccess WindowsDefenderApplicationControlSupplementalPolicyStatuses = 1
	// WindowsDefenderApplicationControlSupplementalPolicyStatusesVTokenError undocumented
	WindowsDefenderApplicationControlSupplementalPolicyStatusesVTokenError WindowsDefenderApplicationControlSupplementalPolicyStatuses = 2
	// WindowsDefenderApplicationControlSupplementalPolicyStatusesVNotAuthorizedByToken undocumented
	WindowsDefenderApplicationControlSupplementalPolicyStatusesVNotAuthorizedByToken WindowsDefenderApplicationControlSupplementalPolicyStatuses = 3
	// WindowsDefenderApplicationControlSupplementalPolicyStatusesVPolicyNotFound undocumented
	WindowsDefenderApplicationControlSupplementalPolicyStatusesVPolicyNotFound WindowsDefenderApplicationControlSupplementalPolicyStatuses = 4
)

// WindowsDefenderApplicationControlSupplementalPolicyStatusesPUnknown returns a pointer to WindowsDefenderApplicationControlSupplementalPolicyStatusesVUnknown
func WindowsDefenderApplicationControlSupplementalPolicyStatusesPUnknown() *WindowsDefenderApplicationControlSupplementalPolicyStatuses {
	v := WindowsDefenderApplicationControlSupplementalPolicyStatusesVUnknown
	return &v
}

// WindowsDefenderApplicationControlSupplementalPolicyStatusesPSuccess returns a pointer to WindowsDefenderApplicationControlSupplementalPolicyStatusesVSuccess
func WindowsDefenderApplicationControlSupplementalPolicyStatusesPSuccess() *WindowsDefenderApplicationControlSupplementalPolicyStatuses {
	v := WindowsDefenderApplicationControlSupplementalPolicyStatusesVSuccess
	return &v
}

// WindowsDefenderApplicationControlSupplementalPolicyStatusesPTokenError returns a pointer to WindowsDefenderApplicationControlSupplementalPolicyStatusesVTokenError
func WindowsDefenderApplicationControlSupplementalPolicyStatusesPTokenError() *WindowsDefenderApplicationControlSupplementalPolicyStatuses {
	v := WindowsDefenderApplicationControlSupplementalPolicyStatusesVTokenError
	return &v
}

// WindowsDefenderApplicationControlSupplementalPolicyStatusesPNotAuthorizedByToken returns a pointer to WindowsDefenderApplicationControlSupplementalPolicyStatusesVNotAuthorizedByToken
func WindowsDefenderApplicationControlSupplementalPolicyStatusesPNotAuthorizedByToken() *WindowsDefenderApplicationControlSupplementalPolicyStatuses {
	v := WindowsDefenderApplicationControlSupplementalPolicyStatusesVNotAuthorizedByToken
	return &v
}

// WindowsDefenderApplicationControlSupplementalPolicyStatusesPPolicyNotFound returns a pointer to WindowsDefenderApplicationControlSupplementalPolicyStatusesVPolicyNotFound
func WindowsDefenderApplicationControlSupplementalPolicyStatusesPPolicyNotFound() *WindowsDefenderApplicationControlSupplementalPolicyStatuses {
	v := WindowsDefenderApplicationControlSupplementalPolicyStatusesVPolicyNotFound
	return &v
}
