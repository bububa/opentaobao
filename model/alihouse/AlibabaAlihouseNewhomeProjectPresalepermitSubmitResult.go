package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回对象
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult
func GetAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult() *AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult 释放AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult(v *AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectPresalepermitSubmitResult.Put(v)
}
