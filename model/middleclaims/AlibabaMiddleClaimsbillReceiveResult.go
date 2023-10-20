package middleclaims

import (
	"sync"
)

// AlibabaMiddleClaimsbillReceiveResult 结构体
type AlibabaMiddleClaimsbillReceiveResult struct {
	// 系统调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否重复
	Repeated bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolAlibabaMiddleClaimsbillReceiveResult = sync.Pool{
	New: func() any {
		return new(AlibabaMiddleClaimsbillReceiveResult)
	},
}

// GetAlibabaMiddleClaimsbillReceiveResult() 从对象池中获取AlibabaMiddleClaimsbillReceiveResult
func GetAlibabaMiddleClaimsbillReceiveResult() *AlibabaMiddleClaimsbillReceiveResult {
	return poolAlibabaMiddleClaimsbillReceiveResult.Get().(*AlibabaMiddleClaimsbillReceiveResult)
}

// ReleaseAlibabaMiddleClaimsbillReceiveResult 释放AlibabaMiddleClaimsbillReceiveResult
func ReleaseAlibabaMiddleClaimsbillReceiveResult(v *AlibabaMiddleClaimsbillReceiveResult) {
	v.Success = false
	v.Data = false
	v.Repeated = false
	v.Retry = false
	poolAlibabaMiddleClaimsbillReceiveResult.Put(v)
}
