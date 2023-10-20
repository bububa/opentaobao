package baichuan

import (
	"sync"
)

// ErrorCode 结构体
type ErrorCode struct {
	// 详细错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
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
	v.Message = ""
	v.Code = ""
	poolErrorCode.Put(v)
}
