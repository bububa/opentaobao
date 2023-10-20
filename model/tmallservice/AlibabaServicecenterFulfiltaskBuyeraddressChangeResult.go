package tmallservice

import (
	"sync"
)

// AlibabaServicecenterFulfiltaskBuyeraddressChangeResult 结构体
type AlibabaServicecenterFulfiltaskBuyeraddressChangeResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaServicecenterFulfiltaskBuyeraddressChangeResult = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterFulfiltaskBuyeraddressChangeResult)
	},
}

// GetAlibabaServicecenterFulfiltaskBuyeraddressChangeResult() 从对象池中获取AlibabaServicecenterFulfiltaskBuyeraddressChangeResult
func GetAlibabaServicecenterFulfiltaskBuyeraddressChangeResult() *AlibabaServicecenterFulfiltaskBuyeraddressChangeResult {
	return poolAlibabaServicecenterFulfiltaskBuyeraddressChangeResult.Get().(*AlibabaServicecenterFulfiltaskBuyeraddressChangeResult)
}

// ReleaseAlibabaServicecenterFulfiltaskBuyeraddressChangeResult 释放AlibabaServicecenterFulfiltaskBuyeraddressChangeResult
func ReleaseAlibabaServicecenterFulfiltaskBuyeraddressChangeResult(v *AlibabaServicecenterFulfiltaskBuyeraddressChangeResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaServicecenterFulfiltaskBuyeraddressChangeResult.Put(v)
}
