package idleisv

import (
	"sync"
)

// AlibabaIdleIsvUserQueryTopResult 结构体
type AlibabaIdleIsvUserQueryTopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *IdleUserApiDo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvUserQueryTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvUserQueryTopResult)
	},
}

// GetAlibabaIdleIsvUserQueryTopResult() 从对象池中获取AlibabaIdleIsvUserQueryTopResult
func GetAlibabaIdleIsvUserQueryTopResult() *AlibabaIdleIsvUserQueryTopResult {
	return poolAlibabaIdleIsvUserQueryTopResult.Get().(*AlibabaIdleIsvUserQueryTopResult)
}

// ReleaseAlibabaIdleIsvUserQueryTopResult 释放AlibabaIdleIsvUserQueryTopResult
func ReleaseAlibabaIdleIsvUserQueryTopResult(v *AlibabaIdleIsvUserQueryTopResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolAlibabaIdleIsvUserQueryTopResult.Put(v)
}
