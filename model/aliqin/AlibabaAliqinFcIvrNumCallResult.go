package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIvrNumCallResult 结构体
type AlibabaAliqinFcIvrNumCallResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIvrNumCallResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIvrNumCallResult)
	},
}

// GetAlibabaAliqinFcIvrNumCallResult() 从对象池中获取AlibabaAliqinFcIvrNumCallResult
func GetAlibabaAliqinFcIvrNumCallResult() *AlibabaAliqinFcIvrNumCallResult {
	return poolAlibabaAliqinFcIvrNumCallResult.Get().(*AlibabaAliqinFcIvrNumCallResult)
}

// ReleaseAlibabaAliqinFcIvrNumCallResult 释放AlibabaAliqinFcIvrNumCallResult
func ReleaseAlibabaAliqinFcIvrNumCallResult(v *AlibabaAliqinFcIvrNumCallResult) {
	v.Model = ""
	v.ErrCode = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcIvrNumCallResult.Put(v)
}
