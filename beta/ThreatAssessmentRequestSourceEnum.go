// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ThreatAssessmentRequestSource undocumented
type ThreatAssessmentRequestSource int

const (
	// ThreatAssessmentRequestSourceVUndefined undocumented
	ThreatAssessmentRequestSourceVUndefined ThreatAssessmentRequestSource = 0
	// ThreatAssessmentRequestSourceVUser undocumented
	ThreatAssessmentRequestSourceVUser ThreatAssessmentRequestSource = 1
	// ThreatAssessmentRequestSourceVAdministrator undocumented
	ThreatAssessmentRequestSourceVAdministrator ThreatAssessmentRequestSource = 2
)

// ThreatAssessmentRequestSourcePUndefined returns a pointer to ThreatAssessmentRequestSourceVUndefined
func ThreatAssessmentRequestSourcePUndefined() *ThreatAssessmentRequestSource {
	v := ThreatAssessmentRequestSourceVUndefined
	return &v
}

// ThreatAssessmentRequestSourcePUser returns a pointer to ThreatAssessmentRequestSourceVUser
func ThreatAssessmentRequestSourcePUser() *ThreatAssessmentRequestSource {
	v := ThreatAssessmentRequestSourceVUser
	return &v
}

// ThreatAssessmentRequestSourcePAdministrator returns a pointer to ThreatAssessmentRequestSourceVAdministrator
func ThreatAssessmentRequestSourcePAdministrator() *ThreatAssessmentRequestSource {
	v := ThreatAssessmentRequestSourceVAdministrator
	return &v
}
