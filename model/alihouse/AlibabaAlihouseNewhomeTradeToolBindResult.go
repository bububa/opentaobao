package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeTradeToolBindResult 结构体
type AlibabaAlihouseNewhomeTradeToolBindResult struct {
	// data
	Data []TradeToolBindResultDto `json:"data,omitempty" xml:"data>trade_tool_bind_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeTradeToolBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeToolBindResult)
	},
}

// GetAlibabaAlihouseNewhomeTradeToolBindResult() 从对象池中获取AlibabaAlihouseNewhomeTradeToolBindResult
func GetAlibabaAlihouseNewhomeTradeToolBindResult() *AlibabaAlihouseNewhomeTradeToolBindResult {
	return poolAlibabaAlihouseNewhomeTradeToolBindResult.Get().(*AlibabaAlihouseNewhomeTradeToolBindResult)
}

// ReleaseAlibabaAlihouseNewhomeTradeToolBindResult 释放AlibabaAlihouseNewhomeTradeToolBindResult
func ReleaseAlibabaAlihouseNewhomeTradeToolBindResult(v *AlibabaAlihouseNewhomeTradeToolBindResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeTradeToolBindResult.Put(v)
}
