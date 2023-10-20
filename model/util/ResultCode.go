package util

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// 错误码描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误妈code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolResultCode = sync.Pool{
	New: func() any {
		return new(ResultCode)
	},
}

// GetResultCode() 从对象池中获取ResultCode
func GetResultCode() *ResultCode {
	return poolResultCode.Get().(*ResultCode)
}

// ReleaseResultCode 释放ResultCode
func ReleaseResultCode(v *ResultCode) {
	v.Message = ""
	v.Code = 0
	poolResultCode.Put(v)
}
