package wdk

import (
	"sync"
)

// AlibabaWdkMerchantItemQueryResult 结构体
type AlibabaWdkMerchantItemQueryResult struct {
	// errorCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errorDesc
	ErrDesc string `json:"err_desc,omitempty" xml:"err_desc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

var poolAlibabaWdkMerchantItemQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantItemQueryResult)
	},
}

// GetAlibabaWdkMerchantItemQueryResult() 从对象池中获取AlibabaWdkMerchantItemQueryResult
func GetAlibabaWdkMerchantItemQueryResult() *AlibabaWdkMerchantItemQueryResult {
	return poolAlibabaWdkMerchantItemQueryResult.Get().(*AlibabaWdkMerchantItemQueryResult)
}

// ReleaseAlibabaWdkMerchantItemQueryResult 释放AlibabaWdkMerchantItemQueryResult
func ReleaseAlibabaWdkMerchantItemQueryResult(v *AlibabaWdkMerchantItemQueryResult) {
	v.ErrCode = ""
	v.ErrDesc = ""
	v.Result = ""
	v.Suc = false
	poolAlibabaWdkMerchantItemQueryResult.Put(v)
}
