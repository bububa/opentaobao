package cainiaohandover

import (
	"sync"
)

// ErrorInfo 结构体
type ErrorInfo struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
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
	v.ErrorMsg = ""
	poolErrorInfo.Put(v)
}
