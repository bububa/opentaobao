package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeActivityDepositUnbindResult 结构体
type AlibabaAlihouseNewhomeActivityDepositUnbindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否保存成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeActivityDepositUnbindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivityDepositUnbindResult)
	},
}

// GetAlibabaAlihouseNewhomeActivityDepositUnbindResult() 从对象池中获取AlibabaAlihouseNewhomeActivityDepositUnbindResult
func GetAlibabaAlihouseNewhomeActivityDepositUnbindResult() *AlibabaAlihouseNewhomeActivityDepositUnbindResult {
	return poolAlibabaAlihouseNewhomeActivityDepositUnbindResult.Get().(*AlibabaAlihouseNewhomeActivityDepositUnbindResult)
}

// ReleaseAlibabaAlihouseNewhomeActivityDepositUnbindResult 释放AlibabaAlihouseNewhomeActivityDepositUnbindResult
func ReleaseAlibabaAlihouseNewhomeActivityDepositUnbindResult(v *AlibabaAlihouseNewhomeActivityDepositUnbindResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeActivityDepositUnbindResult.Put(v)
}
