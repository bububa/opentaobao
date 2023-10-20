package wdk

import (
	"sync"
)

// AlibabaWdkWholesaleInboundorderCommitApiResult 结构体
type AlibabaWdkWholesaleInboundorderCommitApiResult struct {
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkWholesaleInboundorderCommitApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkWholesaleInboundorderCommitApiResult)
	},
}

// GetAlibabaWdkWholesaleInboundorderCommitApiResult() 从对象池中获取AlibabaWdkWholesaleInboundorderCommitApiResult
func GetAlibabaWdkWholesaleInboundorderCommitApiResult() *AlibabaWdkWholesaleInboundorderCommitApiResult {
	return poolAlibabaWdkWholesaleInboundorderCommitApiResult.Get().(*AlibabaWdkWholesaleInboundorderCommitApiResult)
}

// ReleaseAlibabaWdkWholesaleInboundorderCommitApiResult 释放AlibabaWdkWholesaleInboundorderCommitApiResult
func ReleaseAlibabaWdkWholesaleInboundorderCommitApiResult(v *AlibabaWdkWholesaleInboundorderCommitApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaWdkWholesaleInboundorderCommitApiResult.Put(v)
}
