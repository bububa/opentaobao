package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeShopconfigDetailSubmitResult 结构体
type AlibabaAlihouseNewhomeShopconfigDetailSubmitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeShopconfigDetailSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopconfigDetailSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeShopconfigDetailSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeShopconfigDetailSubmitResult
func GetAlibabaAlihouseNewhomeShopconfigDetailSubmitResult() *AlibabaAlihouseNewhomeShopconfigDetailSubmitResult {
	return poolAlibabaAlihouseNewhomeShopconfigDetailSubmitResult.Get().(*AlibabaAlihouseNewhomeShopconfigDetailSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeShopconfigDetailSubmitResult 释放AlibabaAlihouseNewhomeShopconfigDetailSubmitResult
func ReleaseAlibabaAlihouseNewhomeShopconfigDetailSubmitResult(v *AlibabaAlihouseNewhomeShopconfigDetailSubmitResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeShopconfigDetailSubmitResult.Put(v)
}
