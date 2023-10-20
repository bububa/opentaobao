package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotRechargeCardResult 结构体
type AlibabaAliqinFcIotRechargeCardResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotRechargeCardResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotRechargeCardResult)
	},
}

// GetAlibabaAliqinFcIotRechargeCardResult() 从对象池中获取AlibabaAliqinFcIotRechargeCardResult
func GetAlibabaAliqinFcIotRechargeCardResult() *AlibabaAliqinFcIotRechargeCardResult {
	return poolAlibabaAliqinFcIotRechargeCardResult.Get().(*AlibabaAliqinFcIotRechargeCardResult)
}

// ReleaseAlibabaAliqinFcIotRechargeCardResult 释放AlibabaAliqinFcIotRechargeCardResult
func ReleaseAlibabaAliqinFcIotRechargeCardResult(v *AlibabaAliqinFcIotRechargeCardResult) {
	v.Code = ""
	v.Msg = ""
	v.Model = false
	v.Success = false
	poolAlibabaAliqinFcIotRechargeCardResult.Put(v)
}
