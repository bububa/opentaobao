package wdk

import (
	"sync"
)

// AlibabaWdkWholesaleOrderCommitApiResult 结构体
type AlibabaWdkWholesaleOrderCommitApiResult struct {
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkWholesaleOrderCommitApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkWholesaleOrderCommitApiResult)
	},
}

// GetAlibabaWdkWholesaleOrderCommitApiResult() 从对象池中获取AlibabaWdkWholesaleOrderCommitApiResult
func GetAlibabaWdkWholesaleOrderCommitApiResult() *AlibabaWdkWholesaleOrderCommitApiResult {
	return poolAlibabaWdkWholesaleOrderCommitApiResult.Get().(*AlibabaWdkWholesaleOrderCommitApiResult)
}

// ReleaseAlibabaWdkWholesaleOrderCommitApiResult 释放AlibabaWdkWholesaleOrderCommitApiResult
func ReleaseAlibabaWdkWholesaleOrderCommitApiResult(v *AlibabaWdkWholesaleOrderCommitApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkWholesaleOrderCommitApiResult.Put(v)
}
