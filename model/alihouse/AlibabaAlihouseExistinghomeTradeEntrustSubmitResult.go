package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeTradeEntrustSubmitResult 结构体
type AlibabaAlihouseExistinghomeTradeEntrustSubmitResult struct {
	// 消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeTradeEntrustSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeTradeEntrustSubmitResult)
	},
}

// GetAlibabaAlihouseExistinghomeTradeEntrustSubmitResult() 从对象池中获取AlibabaAlihouseExistinghomeTradeEntrustSubmitResult
func GetAlibabaAlihouseExistinghomeTradeEntrustSubmitResult() *AlibabaAlihouseExistinghomeTradeEntrustSubmitResult {
	return poolAlibabaAlihouseExistinghomeTradeEntrustSubmitResult.Get().(*AlibabaAlihouseExistinghomeTradeEntrustSubmitResult)
}

// ReleaseAlibabaAlihouseExistinghomeTradeEntrustSubmitResult 释放AlibabaAlihouseExistinghomeTradeEntrustSubmitResult
func ReleaseAlibabaAlihouseExistinghomeTradeEntrustSubmitResult(v *AlibabaAlihouseExistinghomeTradeEntrustSubmitResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = false
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeTradeEntrustSubmitResult.Put(v)
}
