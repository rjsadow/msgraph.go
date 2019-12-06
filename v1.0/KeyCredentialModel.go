// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// KeyCredential undocumented
type KeyCredential struct {
	// Object is the base model of KeyCredential
	Object
	// CustomKeyIdentifier undocumented
	CustomKeyIdentifier *Binary `json:"customKeyIdentifier,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Usage undocumented
	Usage *string `json:"usage,omitempty"`
	// Key undocumented
	Key *Binary `json:"key,omitempty"`
}
