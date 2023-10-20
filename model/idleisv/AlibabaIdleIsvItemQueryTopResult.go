package idleisv

import (
	"sync"
)

// AlibabaIdleIsvItemQueryTopResult 结构体
type AlibabaIdleIsvItemQueryTopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *IdleItemApiDo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvItemQueryTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemQueryTopResult)
	},
}

// GetAlibabaIdleIsvItemQueryTopResult() 从对象池中获取AlibabaIdleIsvItemQueryTopResult
func GetAlibabaIdleIsvItemQueryTopResult() *AlibabaIdleIsvItemQueryTopResult {
	return poolAlibabaIdleIsvItemQueryTopResult.Get().(*AlibabaIdleIsvItemQueryTopResult)
}

// ReleaseAlibabaIdleIsvItemQueryTopResult 释放AlibabaIdleIsvItemQueryTopResult
func ReleaseAlibabaIdleIsvItemQueryTopResult(v *AlibabaIdleIsvItemQueryTopResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolAlibabaIdleIsvItemQueryTopResult.Put(v)
}
