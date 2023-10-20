package idleisv

import (
	"sync"
)

// AlibabaIdleIsvOrderAdjustpriceResult 结构体
type AlibabaIdleIsvOrderAdjustpriceResult struct {
	// 修改成功返回标示
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaIdleIsvOrderAdjustpriceResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvOrderAdjustpriceResult)
	},
}

// GetAlibabaIdleIsvOrderAdjustpriceResult() 从对象池中获取AlibabaIdleIsvOrderAdjustpriceResult
func GetAlibabaIdleIsvOrderAdjustpriceResult() *AlibabaIdleIsvOrderAdjustpriceResult {
	return poolAlibabaIdleIsvOrderAdjustpriceResult.Get().(*AlibabaIdleIsvOrderAdjustpriceResult)
}

// ReleaseAlibabaIdleIsvOrderAdjustpriceResult 释放AlibabaIdleIsvOrderAdjustpriceResult
func ReleaseAlibabaIdleIsvOrderAdjustpriceResult(v *AlibabaIdleIsvOrderAdjustpriceResult) {
	v.Data = false
	poolAlibabaIdleIsvOrderAdjustpriceResult.Put(v)
}
