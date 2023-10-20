package xiamicontent

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// result msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// result code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
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
	v.Msg = ""
	v.Code = 0
	v.Success = false
	poolResultCode.Put(v)
}
