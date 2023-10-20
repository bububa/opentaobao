package wdk

import (
	"sync"
)

// AlibabaWdkSkuCombineskuAddApiResults 结构体
type AlibabaWdkSkuCombineskuAddApiResults struct {
	// 商品列表
	Models []AlibabaWdkSkuCombineskuAddApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_combinesku_add_api_result,omitempty"`
	// 接口调用异常编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 接口调用异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuCombineskuAddApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCombineskuAddApiResults)
	},
}

// GetAlibabaWdkSkuCombineskuAddApiResults() 从对象池中获取AlibabaWdkSkuCombineskuAddApiResults
func GetAlibabaWdkSkuCombineskuAddApiResults() *AlibabaWdkSkuCombineskuAddApiResults {
	return poolAlibabaWdkSkuCombineskuAddApiResults.Get().(*AlibabaWdkSkuCombineskuAddApiResults)
}

// ReleaseAlibabaWdkSkuCombineskuAddApiResults 释放AlibabaWdkSkuCombineskuAddApiResults
func ReleaseAlibabaWdkSkuCombineskuAddApiResults(v *AlibabaWdkSkuCombineskuAddApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuCombineskuAddApiResults.Put(v)
}
