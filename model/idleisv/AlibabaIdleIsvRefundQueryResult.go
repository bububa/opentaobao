package idleisv

import (
	"sync"
)

// AlibabaIdleIsvRefundQueryResult 结构体
type AlibabaIdleIsvRefundQueryResult struct {
	// 退款信息
	Module *AppraiseIsvRefundDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleIsvRefundQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvRefundQueryResult)
	},
}

// GetAlibabaIdleIsvRefundQueryResult() 从对象池中获取AlibabaIdleIsvRefundQueryResult
func GetAlibabaIdleIsvRefundQueryResult() *AlibabaIdleIsvRefundQueryResult {
	return poolAlibabaIdleIsvRefundQueryResult.Get().(*AlibabaIdleIsvRefundQueryResult)
}

// ReleaseAlibabaIdleIsvRefundQueryResult 释放AlibabaIdleIsvRefundQueryResult
func ReleaseAlibabaIdleIsvRefundQueryResult(v *AlibabaIdleIsvRefundQueryResult) {
	v.Module = nil
	v.Success = false
	poolAlibabaIdleIsvRefundQueryResult.Put(v)
}
