package wdk

import (
	"sync"
)

// AlibabaWdkTradeOrderBalanceBillQueryApiResult 结构体
type AlibabaWdkTradeOrderBalanceBillQueryApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// model
	Model *OrderBalanceBillResponseDo `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkTradeOrderBalanceBillQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeOrderBalanceBillQueryApiResult)
	},
}

// GetAlibabaWdkTradeOrderBalanceBillQueryApiResult() 从对象池中获取AlibabaWdkTradeOrderBalanceBillQueryApiResult
func GetAlibabaWdkTradeOrderBalanceBillQueryApiResult() *AlibabaWdkTradeOrderBalanceBillQueryApiResult {
	return poolAlibabaWdkTradeOrderBalanceBillQueryApiResult.Get().(*AlibabaWdkTradeOrderBalanceBillQueryApiResult)
}

// ReleaseAlibabaWdkTradeOrderBalanceBillQueryApiResult 释放AlibabaWdkTradeOrderBalanceBillQueryApiResult
func ReleaseAlibabaWdkTradeOrderBalanceBillQueryApiResult(v *AlibabaWdkTradeOrderBalanceBillQueryApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkTradeOrderBalanceBillQueryApiResult.Put(v)
}
