package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectLineResult 结构体
type AlibabaAlihouseNewhomeProjectLineResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectLineResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectLineResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectLineResult() 从对象池中获取AlibabaAlihouseNewhomeProjectLineResult
func GetAlibabaAlihouseNewhomeProjectLineResult() *AlibabaAlihouseNewhomeProjectLineResult {
	return poolAlibabaAlihouseNewhomeProjectLineResult.Get().(*AlibabaAlihouseNewhomeProjectLineResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectLineResult 释放AlibabaAlihouseNewhomeProjectLineResult
func ReleaseAlibabaAlihouseNewhomeProjectLineResult(v *AlibabaAlihouseNewhomeProjectLineResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectLineResult.Put(v)
}
