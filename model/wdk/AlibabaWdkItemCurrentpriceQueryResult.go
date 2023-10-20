package wdk

import (
	"sync"
)

// AlibabaWdkItemCurrentpriceQueryResult 结构体
type AlibabaWdkItemCurrentpriceQueryResult struct {
	// 返回的当前当前商品价格列表
	Models []AlibabaWdkItemCurrentpriceQueryModel `json:"models,omitempty" xml:"models>alibaba_wdk_item_currentprice_query_model,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 异常信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemCurrentpriceQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCurrentpriceQueryResult)
	},
}

// GetAlibabaWdkItemCurrentpriceQueryResult() 从对象池中获取AlibabaWdkItemCurrentpriceQueryResult
func GetAlibabaWdkItemCurrentpriceQueryResult() *AlibabaWdkItemCurrentpriceQueryResult {
	return poolAlibabaWdkItemCurrentpriceQueryResult.Get().(*AlibabaWdkItemCurrentpriceQueryResult)
}

// ReleaseAlibabaWdkItemCurrentpriceQueryResult 释放AlibabaWdkItemCurrentpriceQueryResult
func ReleaseAlibabaWdkItemCurrentpriceQueryResult(v *AlibabaWdkItemCurrentpriceQueryResult) {
	v.Models = v.Models[:0]
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolAlibabaWdkItemCurrentpriceQueryResult.Put(v)
}
