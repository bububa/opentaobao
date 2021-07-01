package fenxiao

// BaseResult 结构体
type BaseResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
