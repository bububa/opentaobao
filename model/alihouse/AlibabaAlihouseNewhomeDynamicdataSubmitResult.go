package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeDynamicdataSubmitResult 结构体
type AlibabaAlihouseNewhomeDynamicdataSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeDynamicdataSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeDynamicdataSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeDynamicdataSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeDynamicdataSubmitResult
func GetAlibabaAlihouseNewhomeDynamicdataSubmitResult() *AlibabaAlihouseNewhomeDynamicdataSubmitResult {
	return poolAlibabaAlihouseNewhomeDynamicdataSubmitResult.Get().(*AlibabaAlihouseNewhomeDynamicdataSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeDynamicdataSubmitResult 释放AlibabaAlihouseNewhomeDynamicdataSubmitResult
func ReleaseAlibabaAlihouseNewhomeDynamicdataSubmitResult(v *AlibabaAlihouseNewhomeDynamicdataSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeDynamicdataSubmitResult.Put(v)
}
