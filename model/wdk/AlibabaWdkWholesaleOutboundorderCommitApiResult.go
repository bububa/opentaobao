package wdk

import (
	"sync"
)

// AlibabaWdkWholesaleOutboundorderCommitApiResult 结构体
type AlibabaWdkWholesaleOutboundorderCommitApiResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkWholesaleOutboundorderCommitApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkWholesaleOutboundorderCommitApiResult)
	},
}

// GetAlibabaWdkWholesaleOutboundorderCommitApiResult() 从对象池中获取AlibabaWdkWholesaleOutboundorderCommitApiResult
func GetAlibabaWdkWholesaleOutboundorderCommitApiResult() *AlibabaWdkWholesaleOutboundorderCommitApiResult {
	return poolAlibabaWdkWholesaleOutboundorderCommitApiResult.Get().(*AlibabaWdkWholesaleOutboundorderCommitApiResult)
}

// ReleaseAlibabaWdkWholesaleOutboundorderCommitApiResult 释放AlibabaWdkWholesaleOutboundorderCommitApiResult
func ReleaseAlibabaWdkWholesaleOutboundorderCommitApiResult(v *AlibabaWdkWholesaleOutboundorderCommitApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkWholesaleOutboundorderCommitApiResult.Put(v)
}
