package wdk

import (
	"sync"
)

// AlibabaWdkMerchantproductEditApiResult 结构体
type AlibabaWdkMerchantproductEditApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 返回结果
	Model *MerchantProductResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkMerchantproductEditApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantproductEditApiResult)
	},
}

// GetAlibabaWdkMerchantproductEditApiResult() 从对象池中获取AlibabaWdkMerchantproductEditApiResult
func GetAlibabaWdkMerchantproductEditApiResult() *AlibabaWdkMerchantproductEditApiResult {
	return poolAlibabaWdkMerchantproductEditApiResult.Get().(*AlibabaWdkMerchantproductEditApiResult)
}

// ReleaseAlibabaWdkMerchantproductEditApiResult 释放AlibabaWdkMerchantproductEditApiResult
func ReleaseAlibabaWdkMerchantproductEditApiResult(v *AlibabaWdkMerchantproductEditApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkMerchantproductEditApiResult.Put(v)
}
