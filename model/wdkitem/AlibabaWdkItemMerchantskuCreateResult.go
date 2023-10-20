package wdkitem

import (
	"sync"
)

// AlibabaWdkItemMerchantskuCreateResult 结构体
type AlibabaWdkItemMerchantskuCreateResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemMerchantskuCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantskuCreateResult)
	},
}

// GetAlibabaWdkItemMerchantskuCreateResult() 从对象池中获取AlibabaWdkItemMerchantskuCreateResult
func GetAlibabaWdkItemMerchantskuCreateResult() *AlibabaWdkItemMerchantskuCreateResult {
	return poolAlibabaWdkItemMerchantskuCreateResult.Get().(*AlibabaWdkItemMerchantskuCreateResult)
}

// ReleaseAlibabaWdkItemMerchantskuCreateResult 释放AlibabaWdkItemMerchantskuCreateResult
func ReleaseAlibabaWdkItemMerchantskuCreateResult(v *AlibabaWdkItemMerchantskuCreateResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Code = ""
	v.Success = false
	poolAlibabaWdkItemMerchantskuCreateResult.Put(v)
}
