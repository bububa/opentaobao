package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMerchantskuUpdateResult 结构体
type AlibabaWdkItemMerchantskuUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMerchantskuUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantskuUpdateResult)
	},
}

// GetAlibabaWdkItemMerchantskuUpdateResult() 从对象池中获取AlibabaWdkItemMerchantskuUpdateResult
func GetAlibabaWdkItemMerchantskuUpdateResult() *AlibabaWdkItemMerchantskuUpdateResult {
	return poolAlibabaWdkItemMerchantskuUpdateResult.Get().(*AlibabaWdkItemMerchantskuUpdateResult)
}

// ReleaseAlibabaWdkItemMerchantskuUpdateResult 释放AlibabaWdkItemMerchantskuUpdateResult
func ReleaseAlibabaWdkItemMerchantskuUpdateResult(v *AlibabaWdkItemMerchantskuUpdateResult) {
	v.Code = ""
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Result = ""
	v.Success = false
	poolAlibabaWdkItemMerchantskuUpdateResult.Put(v)
}
