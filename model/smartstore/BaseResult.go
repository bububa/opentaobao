package smartstore

// BaseResult 结构体
type BaseResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
