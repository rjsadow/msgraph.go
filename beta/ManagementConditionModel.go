// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ManagementCondition Management conditions are events that can be triggered dynamically such as geo-fences, time-fences, and network-fences.
type ManagementCondition struct {
	Entity
	// UniqueName Unique name for the management condition. Used in management condition expressions.
	UniqueName *string `json:"uniqueName,omitempty"`
	// DisplayName The admin defined name of the management condition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The admin defined description of the management condition.
	Description *string `json:"description,omitempty"`
	// CreatedDateTime The time the management condition was created. Generated service side.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedDateTime The time the management condition was last modified. Updated service side.
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// ETag ETag of the management condition. Updated service side.
	ETag *string `json:"eTag,omitempty"`
	// ApplicablePlatforms The applicable platforms for this management condition.
	ApplicablePlatforms []DevicePlatformType `json:"applicablePlatforms,omitempty"`
	// ManagementConditionStatements undocumented
	ManagementConditionStatements []ManagementConditionStatement `json:"managementConditionStatements,omitempty"`
}