package alihouse

import (
	"sync"
)

// AlibabaAlihouseMerchantOpenUpdateResult 结构体
type AlibabaAlihouseMerchantOpenUpdateResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseMerchantOpenUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMerchantOpenUpdateResult)
	},
}

// GetAlibabaAlihouseMerchantOpenUpdateResult() 从对象池中获取AlibabaAlihouseMerchantOpenUpdateResult
func GetAlibabaAlihouseMerchantOpenUpdateResult() *AlibabaAlihouseMerchantOpenUpdateResult {
	return poolAlibabaAlihouseMerchantOpenUpdateResult.Get().(*AlibabaAlihouseMerchantOpenUpdateResult)
}

// ReleaseAlibabaAlihouseMerchantOpenUpdateResult 释放AlibabaAlihouseMerchantOpenUpdateResult
func ReleaseAlibabaAlihouseMerchantOpenUpdateResult(v *AlibabaAlihouseMerchantOpenUpdateResult) {
	v.Code = ""
	v.Msg = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseMerchantOpenUpdateResult.Put(v)
}
