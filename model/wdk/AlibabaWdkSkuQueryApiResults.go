package wdk

import (
	"sync"
)

// AlibabaWdkSkuQueryApiResults 结构体
type AlibabaWdkSkuQueryApiResults struct {
	// 结果集合
	Models []AlibabaWdkSkuQueryApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_query_api_result,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuQueryApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuQueryApiResults)
	},
}

// GetAlibabaWdkSkuQueryApiResults() 从对象池中获取AlibabaWdkSkuQueryApiResults
func GetAlibabaWdkSkuQueryApiResults() *AlibabaWdkSkuQueryApiResults {
	return poolAlibabaWdkSkuQueryApiResults.Get().(*AlibabaWdkSkuQueryApiResults)
}

// ReleaseAlibabaWdkSkuQueryApiResults 释放AlibabaWdkSkuQueryApiResults
func ReleaseAlibabaWdkSkuQueryApiResults(v *AlibabaWdkSkuQueryApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuQueryApiResults.Put(v)
}
