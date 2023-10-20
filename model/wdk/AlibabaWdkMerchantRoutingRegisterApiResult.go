package wdk

import (
	"sync"
)

// AlibabaWdkMerchantRoutingRegisterApiResult 结构体
type AlibabaWdkMerchantRoutingRegisterApiResult struct {
	// 调用接口错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用结果返回成功失败
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkMerchantRoutingRegisterApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantRoutingRegisterApiResult)
	},
}

// GetAlibabaWdkMerchantRoutingRegisterApiResult() 从对象池中获取AlibabaWdkMerchantRoutingRegisterApiResult
func GetAlibabaWdkMerchantRoutingRegisterApiResult() *AlibabaWdkMerchantRoutingRegisterApiResult {
	return poolAlibabaWdkMerchantRoutingRegisterApiResult.Get().(*AlibabaWdkMerchantRoutingRegisterApiResult)
}

// ReleaseAlibabaWdkMerchantRoutingRegisterApiResult 释放AlibabaWdkMerchantRoutingRegisterApiResult
func ReleaseAlibabaWdkMerchantRoutingRegisterApiResult(v *AlibabaWdkMerchantRoutingRegisterApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = false
	v.Success = false
	poolAlibabaWdkMerchantRoutingRegisterApiResult.Put(v)
}
