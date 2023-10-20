package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotModbindResult 结构体
type AlibabaAliqinFcIotModbindResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotModbindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotModbindResult)
	},
}

// GetAlibabaAliqinFcIotModbindResult() 从对象池中获取AlibabaAliqinFcIotModbindResult
func GetAlibabaAliqinFcIotModbindResult() *AlibabaAliqinFcIotModbindResult {
	return poolAlibabaAliqinFcIotModbindResult.Get().(*AlibabaAliqinFcIotModbindResult)
}

// ReleaseAlibabaAliqinFcIotModbindResult 释放AlibabaAliqinFcIotModbindResult
func ReleaseAlibabaAliqinFcIotModbindResult(v *AlibabaAliqinFcIotModbindResult) {
	v.Model = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcIotModbindResult.Put(v)
}
