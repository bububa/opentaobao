package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotCardofferResult 结构体
type AlibabaAliqinFcIotCardofferResult struct {
	// 结果对象
	Models []AlibabaAliqinFcIotCardofferModel `json:"models,omitempty" xml:"models>alibaba_aliqin_fc_iot_cardoffer_model,omitempty"`
	// 1.成功；2.失败
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotCardofferResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotCardofferResult)
	},
}

// GetAlibabaAliqinFcIotCardofferResult() 从对象池中获取AlibabaAliqinFcIotCardofferResult
func GetAlibabaAliqinFcIotCardofferResult() *AlibabaAliqinFcIotCardofferResult {
	return poolAlibabaAliqinFcIotCardofferResult.Get().(*AlibabaAliqinFcIotCardofferResult)
}

// ReleaseAlibabaAliqinFcIotCardofferResult 释放AlibabaAliqinFcIotCardofferResult
func ReleaseAlibabaAliqinFcIotCardofferResult(v *AlibabaAliqinFcIotCardofferResult) {
	v.Models = v.Models[:0]
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcIotCardofferResult.Put(v)
}
