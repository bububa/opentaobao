package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeBaseLabelSubmitResult 结构体
type AlibabaAlihouseNewhomeBaseLabelSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeBaseLabelSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeBaseLabelSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeBaseLabelSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeBaseLabelSubmitResult
func GetAlibabaAlihouseNewhomeBaseLabelSubmitResult() *AlibabaAlihouseNewhomeBaseLabelSubmitResult {
	return poolAlibabaAlihouseNewhomeBaseLabelSubmitResult.Get().(*AlibabaAlihouseNewhomeBaseLabelSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeBaseLabelSubmitResult 释放AlibabaAlihouseNewhomeBaseLabelSubmitResult
func ReleaseAlibabaAlihouseNewhomeBaseLabelSubmitResult(v *AlibabaAlihouseNewhomeBaseLabelSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeBaseLabelSubmitResult.Put(v)
}
