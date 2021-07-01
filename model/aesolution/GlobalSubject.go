package aesolution

// GlobalSubject 结构体
type GlobalSubject struct {
	// locale of the subject
	Locale string `json:"locale,omitempty" xml:"locale,omitempty"`
	// subject
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
}
