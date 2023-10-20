package middleclaims

import (
	"sync"
)

// AlibabaMiddleClaimsacceptReceiveResult 结构体
type AlibabaMiddleClaimsacceptReceiveResult struct {
	// 系统调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否重复
	Repeated bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolAlibabaMiddleClaimsacceptReceiveResult = sync.Pool{
	New: func() any {
		return new(AlibabaMiddleClaimsacceptReceiveResult)
	},
}

// GetAlibabaMiddleClaimsacceptReceiveResult() 从对象池中获取AlibabaMiddleClaimsacceptReceiveResult
func GetAlibabaMiddleClaimsacceptReceiveResult() *AlibabaMiddleClaimsacceptReceiveResult {
	return poolAlibabaMiddleClaimsacceptReceiveResult.Get().(*AlibabaMiddleClaimsacceptReceiveResult)
}

// ReleaseAlibabaMiddleClaimsacceptReceiveResult 释放AlibabaMiddleClaimsacceptReceiveResult
func ReleaseAlibabaMiddleClaimsacceptReceiveResult(v *AlibabaMiddleClaimsacceptReceiveResult) {
	v.Success = false
	v.Data = false
	v.Repeated = false
	v.Retry = false
	poolAlibabaMiddleClaimsacceptReceiveResult.Put(v)
}
