package baichuan

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
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
	v.Code = ""
	poolResultCode.Put(v)
}
