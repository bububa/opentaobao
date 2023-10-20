package wdk

import (
	"sync"
)

// AlibabaWdkSkuCombineskuQueryApiResults 结构体
type AlibabaWdkSkuCombineskuQueryApiResults struct {
	// 商品列表
	Models []AlibabaWdkSkuCombineskuQueryApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_combinesku_query_api_result,omitempty"`
	// 接口调用异常编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 接口调用异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuCombineskuQueryApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCombineskuQueryApiResults)
	},
}

// GetAlibabaWdkSkuCombineskuQueryApiResults() 从对象池中获取AlibabaWdkSkuCombineskuQueryApiResults
func GetAlibabaWdkSkuCombineskuQueryApiResults() *AlibabaWdkSkuCombineskuQueryApiResults {
	return poolAlibabaWdkSkuCombineskuQueryApiResults.Get().(*AlibabaWdkSkuCombineskuQueryApiResults)
}

// ReleaseAlibabaWdkSkuCombineskuQueryApiResults 释放AlibabaWdkSkuCombineskuQueryApiResults
func ReleaseAlibabaWdkSkuCombineskuQueryApiResults(v *AlibabaWdkSkuCombineskuQueryApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuCombineskuQueryApiResults.Put(v)
}
