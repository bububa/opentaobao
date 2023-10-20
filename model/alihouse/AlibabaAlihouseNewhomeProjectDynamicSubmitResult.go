package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectDynamicSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectDynamicSubmitResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回对象
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectDynamicSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectDynamicSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectDynamicSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectDynamicSubmitResult
func GetAlibabaAlihouseNewhomeProjectDynamicSubmitResult() *AlibabaAlihouseNewhomeProjectDynamicSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectDynamicSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectDynamicSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectDynamicSubmitResult 释放AlibabaAlihouseNewhomeProjectDynamicSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectDynamicSubmitResult(v *AlibabaAlihouseNewhomeProjectDynamicSubmitResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeProjectDynamicSubmitResult.Put(v)
}
