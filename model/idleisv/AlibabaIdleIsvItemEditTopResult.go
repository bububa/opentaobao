package idleisv

import (
	"sync"
)

// AlibabaIdleIsvItemEditTopResult 结构体
type AlibabaIdleIsvItemEditTopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *IdleItemApiDo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvItemEditTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvItemEditTopResult)
	},
}

// GetAlibabaIdleIsvItemEditTopResult() 从对象池中获取AlibabaIdleIsvItemEditTopResult
func GetAlibabaIdleIsvItemEditTopResult() *AlibabaIdleIsvItemEditTopResult {
	return poolAlibabaIdleIsvItemEditTopResult.Get().(*AlibabaIdleIsvItemEditTopResult)
}

// ReleaseAlibabaIdleIsvItemEditTopResult 释放AlibabaIdleIsvItemEditTopResult
func ReleaseAlibabaIdleIsvItemEditTopResult(v *AlibabaIdleIsvItemEditTopResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Data = nil
	v.Success = false
	poolAlibabaIdleIsvItemEditTopResult.Put(v)
}
