package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult 结构体
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult() 从对象池中获取AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult
func GetAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult() *AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult {
	return poolAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult.Get().(*AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult 释放AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult
func ReleaseAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult(v *AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectPresalepermitDeleteResult.Put(v)
}
