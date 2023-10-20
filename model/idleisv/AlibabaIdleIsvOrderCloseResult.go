package idleisv

import (
	"sync"
)

// AlibabaIdleIsvOrderCloseResult 结构体
type AlibabaIdleIsvOrderCloseResult struct {
	// 关闭成功返回标识
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaIdleIsvOrderCloseResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvOrderCloseResult)
	},
}

// GetAlibabaIdleIsvOrderCloseResult() 从对象池中获取AlibabaIdleIsvOrderCloseResult
func GetAlibabaIdleIsvOrderCloseResult() *AlibabaIdleIsvOrderCloseResult {
	return poolAlibabaIdleIsvOrderCloseResult.Get().(*AlibabaIdleIsvOrderCloseResult)
}

// ReleaseAlibabaIdleIsvOrderCloseResult 释放AlibabaIdleIsvOrderCloseResult
func ReleaseAlibabaIdleIsvOrderCloseResult(v *AlibabaIdleIsvOrderCloseResult) {
	v.Data = false
	poolAlibabaIdleIsvOrderCloseResult.Put(v)
}
