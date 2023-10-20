package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMerchantskuQueryResult 结构体
type AlibabaWdkItemMerchantskuQueryResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMerchantskuQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantskuQueryResult)
	},
}

// GetAlibabaWdkItemMerchantskuQueryResult() 从对象池中获取AlibabaWdkItemMerchantskuQueryResult
func GetAlibabaWdkItemMerchantskuQueryResult() *AlibabaWdkItemMerchantskuQueryResult {
	return poolAlibabaWdkItemMerchantskuQueryResult.Get().(*AlibabaWdkItemMerchantskuQueryResult)
}

// ReleaseAlibabaWdkItemMerchantskuQueryResult 释放AlibabaWdkItemMerchantskuQueryResult
func ReleaseAlibabaWdkItemMerchantskuQueryResult(v *AlibabaWdkItemMerchantskuQueryResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Result = ""
	v.Code = ""
	v.Success = false
	poolAlibabaWdkItemMerchantskuQueryResult.Put(v)
}
