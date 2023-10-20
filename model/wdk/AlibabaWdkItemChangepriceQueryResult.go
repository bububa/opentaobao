package wdk

import (
	"sync"
)

// AlibabaWdkItemChangepriceQueryResult 结构体
type AlibabaWdkItemChangepriceQueryResult struct {
	// 返回的商品改价单
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// 结果码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 结果文案
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkItemChangepriceQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemChangepriceQueryResult)
	},
}

// GetAlibabaWdkItemChangepriceQueryResult() 从对象池中获取AlibabaWdkItemChangepriceQueryResult
func GetAlibabaWdkItemChangepriceQueryResult() *AlibabaWdkItemChangepriceQueryResult {
	return poolAlibabaWdkItemChangepriceQueryResult.Get().(*AlibabaWdkItemChangepriceQueryResult)
}

// ReleaseAlibabaWdkItemChangepriceQueryResult 释放AlibabaWdkItemChangepriceQueryResult
func ReleaseAlibabaWdkItemChangepriceQueryResult(v *AlibabaWdkItemChangepriceQueryResult) {
	v.Models = v.Models[:0]
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolAlibabaWdkItemChangepriceQueryResult.Put(v)
}
