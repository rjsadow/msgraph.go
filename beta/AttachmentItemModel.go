// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AttachmentItem undocumented
type AttachmentItem struct {
	// Object is the base model of AttachmentItem
	Object
	// AttachmentType undocumented
	AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Size undocumented
	Size *int `json:"size,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// IsInline undocumented
	IsInline *bool `json:"isInline,omitempty"`
}
