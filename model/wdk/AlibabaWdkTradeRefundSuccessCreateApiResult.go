package wdk

import (
	"sync"
)

// AlibabaWdkTradeRefundSuccessCreateApiResult 结构体
type AlibabaWdkTradeRefundSuccessCreateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口状态
	State bool `json:"state,omitempty" xml:"state,omitempty"`
	// 是否成功 true-成功;false-失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolAlibabaWdkTradeRefundSuccessCreateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeRefundSuccessCreateApiResult)
	},
}

// GetAlibabaWdkTradeRefundSuccessCreateApiResult() 从对象池中获取AlibabaWdkTradeRefundSuccessCreateApiResult
func GetAlibabaWdkTradeRefundSuccessCreateApiResult() *AlibabaWdkTradeRefundSuccessCreateApiResult {
	return poolAlibabaWdkTradeRefundSuccessCreateApiResult.Get().(*AlibabaWdkTradeRefundSuccessCreateApiResult)
}

// ReleaseAlibabaWdkTradeRefundSuccessCreateApiResult 释放AlibabaWdkTradeRefundSuccessCreateApiResult
func ReleaseAlibabaWdkTradeRefundSuccessCreateApiResult(v *AlibabaWdkTradeRefundSuccessCreateApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.State = false
	v.Result = false
	poolAlibabaWdkTradeRefundSuccessCreateApiResult.Put(v)
}
