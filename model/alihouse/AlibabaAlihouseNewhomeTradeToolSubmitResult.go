package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeTradeToolSubmitResult 结构体
type AlibabaAlihouseNewhomeTradeToolSubmitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeTradeToolSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeToolSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeTradeToolSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeTradeToolSubmitResult
func GetAlibabaAlihouseNewhomeTradeToolSubmitResult() *AlibabaAlihouseNewhomeTradeToolSubmitResult {
	return poolAlibabaAlihouseNewhomeTradeToolSubmitResult.Get().(*AlibabaAlihouseNewhomeTradeToolSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeTradeToolSubmitResult 释放AlibabaAlihouseNewhomeTradeToolSubmitResult
func ReleaseAlibabaAlihouseNewhomeTradeToolSubmitResult(v *AlibabaAlihouseNewhomeTradeToolSubmitResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeTradeToolSubmitResult.Put(v)
}
