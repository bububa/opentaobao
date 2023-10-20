package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectSubmitResult 结构体
type AlibabaAlihouseNewhomeProjectSubmitResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口返回对象
	Data *EbbasProjectSubmitVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeProjectSubmitResult
func GetAlibabaAlihouseNewhomeProjectSubmitResult() *AlibabaAlihouseNewhomeProjectSubmitResult {
	return poolAlibabaAlihouseNewhomeProjectSubmitResult.Get().(*AlibabaAlihouseNewhomeProjectSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectSubmitResult 释放AlibabaAlihouseNewhomeProjectSubmitResult
func ReleaseAlibabaAlihouseNewhomeProjectSubmitResult(v *AlibabaAlihouseNewhomeProjectSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectSubmitResult.Put(v)
}
