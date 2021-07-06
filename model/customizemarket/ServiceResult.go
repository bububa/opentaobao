package customizemarket

// ServiceResult 结构体
type ServiceResult struct {
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// suc
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}
