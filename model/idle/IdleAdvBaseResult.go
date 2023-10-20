package idle

import (
	"sync"
)

// IdleAdvBaseResult 结构体
type IdleAdvBaseResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误原因描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

var poolIdleAdvBaseResult = sync.Pool{
	New: func() any {
		return new(IdleAdvBaseResult)
	},
}

// GetIdleAdvBaseResult() 从对象池中获取IdleAdvBaseResult
func GetIdleAdvBaseResult() *IdleAdvBaseResult {
	return poolIdleAdvBaseResult.Get().(*IdleAdvBaseResult)
}

// ReleaseIdleAdvBaseResult 释放IdleAdvBaseResult
func ReleaseIdleAdvBaseResult(v *IdleAdvBaseResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	poolIdleAdvBaseResult.Put(v)
}
