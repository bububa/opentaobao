package wdk

import (
	"sync"
)

// AlibabaWdkCouponContractCreateApiResult 结构体
type AlibabaWdkCouponContractCreateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 合同ID
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkCouponContractCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkCouponContractCreateApiResult)
	},
}

// GetAlibabaWdkCouponContractCreateApiResult() 从对象池中获取AlibabaWdkCouponContractCreateApiResult
func GetAlibabaWdkCouponContractCreateApiResult() *AlibabaWdkCouponContractCreateApiResult {
	return poolAlibabaWdkCouponContractCreateApiResult.Get().(*AlibabaWdkCouponContractCreateApiResult)
}

// ReleaseAlibabaWdkCouponContractCreateApiResult 释放AlibabaWdkCouponContractCreateApiResult
func ReleaseAlibabaWdkCouponContractCreateApiResult(v *AlibabaWdkCouponContractCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = 0
	v.Success = false
	poolAlibabaWdkCouponContractCreateApiResult.Put(v)
}
