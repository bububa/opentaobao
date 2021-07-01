package vms

// AppBaseResponse 结构体
type AppBaseResponse struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
