package middleclaims

import (
	"sync"
)

// AlibabaMiddleClaimsresultReceiveResult 结构体
type AlibabaMiddleClaimsresultReceiveResult struct {
	// 系统调用结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否重复
	Repeated bool `json:"repeated,omitempty" xml:"repeated,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

var poolAlibabaMiddleClaimsresultReceiveResult = sync.Pool{
	New: func() any {
		return new(AlibabaMiddleClaimsresultReceiveResult)
	},
}

// GetAlibabaMiddleClaimsresultReceiveResult() 从对象池中获取AlibabaMiddleClaimsresultReceiveResult
func GetAlibabaMiddleClaimsresultReceiveResult() *AlibabaMiddleClaimsresultReceiveResult {
	return poolAlibabaMiddleClaimsresultReceiveResult.Get().(*AlibabaMiddleClaimsresultReceiveResult)
}

// ReleaseAlibabaMiddleClaimsresultReceiveResult 释放AlibabaMiddleClaimsresultReceiveResult
func ReleaseAlibabaMiddleClaimsresultReceiveResult(v *AlibabaMiddleClaimsresultReceiveResult) {
	v.Success = false
	v.Data = false
	v.Repeated = false
	v.Retry = false
	poolAlibabaMiddleClaimsresultReceiveResult.Put(v)
}
