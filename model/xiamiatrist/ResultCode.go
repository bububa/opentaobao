package xiamiatrist

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// 是否成功
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
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
	v.Msg = ""
	v.Code = 0
	poolResultCode.Put(v)
}
