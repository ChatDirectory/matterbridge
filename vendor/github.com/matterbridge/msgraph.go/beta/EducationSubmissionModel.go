// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EducationSubmission undocumented
type EducationSubmission struct {
	// Entity is the base model of EducationSubmission
	Entity
	// Recipient undocumented
	Recipient *EducationSubmissionRecipient `json:"recipient,omitempty"`
	// Status undocumented
	Status *EducationSubmissionStatus `json:"status,omitempty"`
	// SubmittedBy undocumented
	SubmittedBy *IdentitySet `json:"submittedBy,omitempty"`
	// SubmittedDateTime undocumented
	SubmittedDateTime *time.Time `json:"submittedDateTime,omitempty"`
	// UnsubmittedBy undocumented
	UnsubmittedBy *IdentitySet `json:"unsubmittedBy,omitempty"`
	// UnsubmittedDateTime undocumented
	UnsubmittedDateTime *time.Time `json:"unsubmittedDateTime,omitempty"`
	// ReleasedBy undocumented
	ReleasedBy *IdentitySet `json:"releasedBy,omitempty"`
	// ReleasedDateTime undocumented
	ReleasedDateTime *time.Time `json:"releasedDateTime,omitempty"`
	// ReturnedBy undocumented
	ReturnedBy *IdentitySet `json:"returnedBy,omitempty"`
	// ReturnedDateTime undocumented
	ReturnedDateTime *time.Time `json:"returnedDateTime,omitempty"`
	// ResourcesFolderURL undocumented
	ResourcesFolderURL *string `json:"resourcesFolderUrl,omitempty"`
	// Resources undocumented
	Resources []EducationSubmissionResource `json:"resources,omitempty"`
	// SubmittedResources undocumented
	SubmittedResources []EducationSubmissionResource `json:"submittedResources,omitempty"`
	// Outcomes undocumented
	Outcomes []EducationOutcome `json:"outcomes,omitempty"`
}