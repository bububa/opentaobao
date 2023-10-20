package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectAdviserBindResult 结构体
type AlibabaAlihouseNewhomeProjectAdviserBindResult struct {
	// 操作结果
	Data []ProjectAdviserBindResultDto `json:"data,omitempty" xml:"data>project_adviser_bind_result_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectAdviserBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectAdviserBindResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserBindResult() 从对象池中获取AlibabaAlihouseNewhomeProjectAdviserBindResult
func GetAlibabaAlihouseNewhomeProjectAdviserBindResult() *AlibabaAlihouseNewhomeProjectAdviserBindResult {
	return poolAlibabaAlihouseNewhomeProjectAdviserBindResult.Get().(*AlibabaAlihouseNewhomeProjectAdviserBindResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserBindResult 释放AlibabaAlihouseNewhomeProjectAdviserBindResult
func ReleaseAlibabaAlihouseNewhomeProjectAdviserBindResult(v *AlibabaAlihouseNewhomeProjectAdviserBindResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectAdviserBindResult.Put(v)
}
