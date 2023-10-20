package paimai

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
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
	v.Code = ""
	v.Message = ""
	poolResultCode.Put(v)
}
