package alihouse

import (
	"sync"
)

// AlibabaAlihouseMerchantTradeConfigBindResult 结构体
type AlibabaAlihouseMerchantTradeConfigBindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseMerchantTradeConfigBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMerchantTradeConfigBindResult)
	},
}

// GetAlibabaAlihouseMerchantTradeConfigBindResult() 从对象池中获取AlibabaAlihouseMerchantTradeConfigBindResult
func GetAlibabaAlihouseMerchantTradeConfigBindResult() *AlibabaAlihouseMerchantTradeConfigBindResult {
	return poolAlibabaAlihouseMerchantTradeConfigBindResult.Get().(*AlibabaAlihouseMerchantTradeConfigBindResult)
}

// ReleaseAlibabaAlihouseMerchantTradeConfigBindResult 释放AlibabaAlihouseMerchantTradeConfigBindResult
func ReleaseAlibabaAlihouseMerchantTradeConfigBindResult(v *AlibabaAlihouseMerchantTradeConfigBindResult) {
	v.Code = ""
	v.Msg = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseMerchantTradeConfigBindResult.Put(v)
}
