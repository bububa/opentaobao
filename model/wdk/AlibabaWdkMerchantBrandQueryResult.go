package wdk

import (
	"sync"
)

// AlibabaWdkMerchantBrandQueryResult 结构体
type AlibabaWdkMerchantBrandQueryResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// result
	Results string `json:"results,omitempty" xml:"results,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkMerchantBrandQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantBrandQueryResult)
	},
}

// GetAlibabaWdkMerchantBrandQueryResult() 从对象池中获取AlibabaWdkMerchantBrandQueryResult
func GetAlibabaWdkMerchantBrandQueryResult() *AlibabaWdkMerchantBrandQueryResult {
	return poolAlibabaWdkMerchantBrandQueryResult.Get().(*AlibabaWdkMerchantBrandQueryResult)
}

// ReleaseAlibabaWdkMerchantBrandQueryResult 释放AlibabaWdkMerchantBrandQueryResult
func ReleaseAlibabaWdkMerchantBrandQueryResult(v *AlibabaWdkMerchantBrandQueryResult) {
	v.ErrorCode = ""
	v.ErrorDesc = ""
	v.Results = ""
	v.Success = false
	poolAlibabaWdkMerchantBrandQueryResult.Put(v)
}
