// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkMobileAppConfiguration Contains properties, inherited properties and actions for AFW mobile app configurations.
type AndroidForWorkMobileAppConfiguration struct {
	ManagedDeviceMobileAppConfiguration
	// PackageID Android For Work app configuration package id.
	PackageID *string `json:"packageId,omitempty"`
	// PayloadJSON Android For Work app configuration JSON payload.
	PayloadJSON *string `json:"payloadJson,omitempty"`
	// PermissionActions List of Android app permissions and corresponding permission actions.
	PermissionActions []AndroidPermissionAction `json:"permissionActions,omitempty"`
}