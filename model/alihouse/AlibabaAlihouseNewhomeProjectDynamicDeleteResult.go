package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectDynamicDeleteResult 结构体
type AlibabaAlihouseNewhomeProjectDynamicDeleteResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectDynamicDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectDynamicDeleteResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectDynamicDeleteResult() 从对象池中获取AlibabaAlihouseNewhomeProjectDynamicDeleteResult
func GetAlibabaAlihouseNewhomeProjectDynamicDeleteResult() *AlibabaAlihouseNewhomeProjectDynamicDeleteResult {
	return poolAlibabaAlihouseNewhomeProjectDynamicDeleteResult.Get().(*AlibabaAlihouseNewhomeProjectDynamicDeleteResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectDynamicDeleteResult 释放AlibabaAlihouseNewhomeProjectDynamicDeleteResult
func ReleaseAlibabaAlihouseNewhomeProjectDynamicDeleteResult(v *AlibabaAlihouseNewhomeProjectDynamicDeleteResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectDynamicDeleteResult.Put(v)
}
