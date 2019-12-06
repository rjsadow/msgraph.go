// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Application undocumented
type Application struct {
	// DirectoryObject is the base model of Application
	DirectoryObject
	// AddIns undocumented
	AddIns []AddIn `json:"addIns,omitempty"`
	// API undocumented
	API *APIApplication `json:"api,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// ApplicationTemplateID undocumented
	ApplicationTemplateID *string `json:"applicationTemplateId,omitempty"`
	// AppRoles undocumented
	AppRoles []AppRole `json:"appRoles,omitempty"`
	// IsFallbackPublicClient undocumented
	IsFallbackPublicClient *bool `json:"isFallbackPublicClient,omitempty"`
	// IdentifierUris undocumented
	IdentifierUris []string `json:"identifierUris,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// PublicClient undocumented
	PublicClient *PublicClientApplication `json:"publicClient,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GroupMembershipClaims undocumented
	GroupMembershipClaims *string `json:"groupMembershipClaims,omitempty"`
	// Info undocumented
	Info *InformationalURL `json:"info,omitempty"`
	// IsDeviceOnlyAuthSupported undocumented
	IsDeviceOnlyAuthSupported *bool `json:"isDeviceOnlyAuthSupported,omitempty"`
	// KeyCredentials undocumented
	KeyCredentials []KeyCredential `json:"keyCredentials,omitempty"`
	// Logo undocumented
	Logo *Stream `json:"logo,omitempty"`
	// Oauth2RequirePostResponse undocumented
	Oauth2RequirePostResponse *bool `json:"oauth2RequirePostResponse,omitempty"`
	// OptionalClaims undocumented
	OptionalClaims *OptionalClaims `json:"optionalClaims,omitempty"`
	// ParentalControlSettings undocumented
	ParentalControlSettings *ParentalControlSettings `json:"parentalControlSettings,omitempty"`
	// PasswordCredentials undocumented
	PasswordCredentials []PasswordCredential `json:"passwordCredentials,omitempty"`
	// PublisherDomain undocumented
	PublisherDomain *string `json:"publisherDomain,omitempty"`
	// RequiredResourceAccess undocumented
	RequiredResourceAccess []RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`
	// SignInAudience undocumented
	SignInAudience *string `json:"signInAudience,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// TokenEncryptionKeyID undocumented
	TokenEncryptionKeyID *UUID `json:"tokenEncryptionKeyId,omitempty"`
	// Web undocumented
	Web *WebApplication `json:"web,omitempty"`
	// ExtensionProperties undocumented
	ExtensionProperties []ExtensionProperty `json:"extensionProperties,omitempty"`
	// CreatedOnBehalfOf undocumented
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
	// Owners undocumented
	Owners []DirectoryObject `json:"owners,omitempty"`
}
