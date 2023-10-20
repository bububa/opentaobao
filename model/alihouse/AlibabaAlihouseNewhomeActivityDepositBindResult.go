package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeActivityDepositBindResult 结构体
type AlibabaAlihouseNewhomeActivityDepositBindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否保存成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeActivityDepositBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivityDepositBindResult)
	},
}

// GetAlibabaAlihouseNewhomeActivityDepositBindResult() 从对象池中获取AlibabaAlihouseNewhomeActivityDepositBindResult
func GetAlibabaAlihouseNewhomeActivityDepositBindResult() *AlibabaAlihouseNewhomeActivityDepositBindResult {
	return poolAlibabaAlihouseNewhomeActivityDepositBindResult.Get().(*AlibabaAlihouseNewhomeActivityDepositBindResult)
}

// ReleaseAlibabaAlihouseNewhomeActivityDepositBindResult 释放AlibabaAlihouseNewhomeActivityDepositBindResult
func ReleaseAlibabaAlihouseNewhomeActivityDepositBindResult(v *AlibabaAlihouseNewhomeActivityDepositBindResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeActivityDepositBindResult.Put(v)
}
