package icbuassurance

// AssuranceAccountResult 结构体
type AssuranceAccountResult struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// value
	Value *AssuranceFlag `json:"value,omitempty" xml:"value,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
