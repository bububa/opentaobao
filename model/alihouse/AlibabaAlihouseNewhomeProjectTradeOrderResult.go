package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectTradeOrderResult 结构体
type AlibabaAlihouseNewhomeProjectTradeOrderResult struct {
	// 成功的楼盘Id
	Data []string `json:"data,omitempty" xml:"data>string,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectTradeOrderResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectTradeOrderResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectTradeOrderResult() 从对象池中获取AlibabaAlihouseNewhomeProjectTradeOrderResult
func GetAlibabaAlihouseNewhomeProjectTradeOrderResult() *AlibabaAlihouseNewhomeProjectTradeOrderResult {
	return poolAlibabaAlihouseNewhomeProjectTradeOrderResult.Get().(*AlibabaAlihouseNewhomeProjectTradeOrderResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectTradeOrderResult 释放AlibabaAlihouseNewhomeProjectTradeOrderResult
func ReleaseAlibabaAlihouseNewhomeProjectTradeOrderResult(v *AlibabaAlihouseNewhomeProjectTradeOrderResult) {
	v.Data = v.Data[:0]
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectTradeOrderResult.Put(v)
}
