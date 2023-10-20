package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectPhoneSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectPhoneSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectPhoneSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectPhoneSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectPhoneSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectPhoneSubmitResult
func GetAlibabaAlihouseNewhomeProjectPhoneSubmitResult() *AlibabaAlihouseNewhomeProjectPhoneSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectPhoneSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectPhoneSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectPhoneSubmitResult 释放AlibabaAlihouseNewhomeProjectPhoneSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectPhoneSubmitResult(v *AlibabaAlihouseNewhomeProjectPhoneSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectPhoneSubmitResult.Put(v)
}
