package media

import (
	"sync"
)

// ResultCode 结构体
type ResultCode struct {
	// 错误描述
	Info string `json:"info,omitempty" xml:"info,omitempty"`
	// 错误代码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
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
	v.Info = ""
	v.ErrorCode = 0
	poolResultCode.Put(v)
}
