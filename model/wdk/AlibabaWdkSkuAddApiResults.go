package wdk

import (
	"sync"
)

// AlibabaWdkSkuAddApiResults 结构体
type AlibabaWdkSkuAddApiResults struct {
	// models
	Models []AlibabaWdkSkuAddApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_add_api_result,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口返回成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuAddApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuAddApiResults)
	},
}

// GetAlibabaWdkSkuAddApiResults() 从对象池中获取AlibabaWdkSkuAddApiResults
func GetAlibabaWdkSkuAddApiResults() *AlibabaWdkSkuAddApiResults {
	return poolAlibabaWdkSkuAddApiResults.Get().(*AlibabaWdkSkuAddApiResults)
}

// ReleaseAlibabaWdkSkuAddApiResults 释放AlibabaWdkSkuAddApiResults
func ReleaseAlibabaWdkSkuAddApiResults(v *AlibabaWdkSkuAddApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuAddApiResults.Put(v)
}
