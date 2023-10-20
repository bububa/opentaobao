package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectAdviserSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectAdviserSubmitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectAdviserSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectAdviserSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectAdviserSubmitResult
func GetAlibabaAlihouseNewhomeProjectAdviserSubmitResult() *AlibabaAlihouseNewhomeProjectAdviserSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectAdviserSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectAdviserSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserSubmitResult 释放AlibabaAlihouseNewhomeProjectAdviserSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectAdviserSubmitResult(v *AlibabaAlihouseNewhomeProjectAdviserSubmitResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeProjectAdviserSubmitResult.Put(v)
}
