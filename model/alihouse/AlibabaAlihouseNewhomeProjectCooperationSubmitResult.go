package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectCooperationSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectCooperationSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectCooperationSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectCooperationSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectCooperationSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectCooperationSubmitResult
func GetAlibabaAlihouseNewhomeProjectCooperationSubmitResult() *AlibabaAlihouseNewhomeProjectCooperationSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectCooperationSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectCooperationSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectCooperationSubmitResult 释放AlibabaAlihouseNewhomeProjectCooperationSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectCooperationSubmitResult(v *AlibabaAlihouseNewhomeProjectCooperationSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectCooperationSubmitResult.Put(v)
}
