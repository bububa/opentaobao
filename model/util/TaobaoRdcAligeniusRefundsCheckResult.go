package util

// TaobaordcaligeniusrefundscheckResult 结构体
type TaobaordcaligeniusrefundscheckResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorInfo
	ErrorInfo string `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
