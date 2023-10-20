package wdk

import (
	"sync"
)

// AlibabaWdkMerchantStoreitemQueryResult 结构体
type AlibabaWdkMerchantStoreitemQueryResult struct {
	// errorCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errorDesc
	ErrDesc string `json:"err_desc,omitempty" xml:"err_desc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// success
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

var poolAlibabaWdkMerchantStoreitemQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMerchantStoreitemQueryResult)
	},
}

// GetAlibabaWdkMerchantStoreitemQueryResult() 从对象池中获取AlibabaWdkMerchantStoreitemQueryResult
func GetAlibabaWdkMerchantStoreitemQueryResult() *AlibabaWdkMerchantStoreitemQueryResult {
	return poolAlibabaWdkMerchantStoreitemQueryResult.Get().(*AlibabaWdkMerchantStoreitemQueryResult)
}

// ReleaseAlibabaWdkMerchantStoreitemQueryResult 释放AlibabaWdkMerchantStoreitemQueryResult
func ReleaseAlibabaWdkMerchantStoreitemQueryResult(v *AlibabaWdkMerchantStoreitemQueryResult) {
	v.ErrCode = ""
	v.ErrDesc = ""
	v.Result = ""
	v.Suc = false
	poolAlibabaWdkMerchantStoreitemQueryResult.Put(v)
}
