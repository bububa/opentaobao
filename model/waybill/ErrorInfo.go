package waybill

import (
	"sync"
)

// ErrorInfo 结构体
type ErrorInfo struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
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
	v.ErrorMessage = ""
	v.ErrorCode = ""
	poolErrorInfo.Put(v)
}
