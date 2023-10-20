package inventory

import (
	"sync"
)

// ErrorInfo 结构体
type ErrorInfo struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// subErrorCode
	SubErrorCode string `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

var poolErrorInfo = sync.Pool{
	New: func() any {
		return new(ErrorInfo)
	},
}

// GetErrorInfo() 从对象池中获取ErrorInfo
func GetErrorInfo() *ErrorInfo {
	return poolErrorInfo.Get().(*ErrorInfo)
}

// ReleaseErrorInfo 释放ErrorInfo
func ReleaseErrorInfo(v *ErrorInfo) {
	v.ErrorCode = ""
	v.SubErrorCode = ""
	v.ErrorMessage = ""
	poolErrorInfo.Put(v)
}
