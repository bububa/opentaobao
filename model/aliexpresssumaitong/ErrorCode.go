package aliexpresssumaitong

import (
	"sync"
)

// ErrorCode 结构体
type ErrorCode struct {
	// 错误展示信息
	DisplayMessage string `json:"display_message,omitempty" xml:"display_message,omitempty"`
	// 错误码key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 错误详情
	LogMessage string `json:"log_message,omitempty" xml:"log_message,omitempty"`
}

var poolErrorCode = sync.Pool{
	New: func() any {
		return new(ErrorCode)
	},
}

// GetErrorCode() 从对象池中获取ErrorCode
func GetErrorCode() *ErrorCode {
	return poolErrorCode.Get().(*ErrorCode)
}

// ReleaseErrorCode 释放ErrorCode
func ReleaseErrorCode(v *ErrorCode) {
	v.DisplayMessage = ""
	v.Key = ""
	v.LogMessage = ""
	poolErrorCode.Put(v)
}
