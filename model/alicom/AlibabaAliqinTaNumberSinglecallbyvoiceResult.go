package alicom

import (
	"sync"
)

// AlibabaAliqinTaNumberSinglecallbyvoiceResult 结构体
type AlibabaAliqinTaNumberSinglecallbyvoiceResult struct {
	// 结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 系统自动生成
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 系统自动生成
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功，失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinTaNumberSinglecallbyvoiceResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaNumberSinglecallbyvoiceResult)
	},
}

// GetAlibabaAliqinTaNumberSinglecallbyvoiceResult() 从对象池中获取AlibabaAliqinTaNumberSinglecallbyvoiceResult
func GetAlibabaAliqinTaNumberSinglecallbyvoiceResult() *AlibabaAliqinTaNumberSinglecallbyvoiceResult {
	return poolAlibabaAliqinTaNumberSinglecallbyvoiceResult.Get().(*AlibabaAliqinTaNumberSinglecallbyvoiceResult)
}

// ReleaseAlibabaAliqinTaNumberSinglecallbyvoiceResult 释放AlibabaAliqinTaNumberSinglecallbyvoiceResult
func ReleaseAlibabaAliqinTaNumberSinglecallbyvoiceResult(v *AlibabaAliqinTaNumberSinglecallbyvoiceResult) {
	v.Model = ""
	v.Msg = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAliqinTaNumberSinglecallbyvoiceResult.Put(v)
}
