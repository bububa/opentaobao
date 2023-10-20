package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult 结构体
type AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult
func GetAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult() *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult {
	return poolAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult.Get().(*AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult 释放AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult
func ReleaseAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult(v *AlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeShopcityconfigDetailSubmitResult.Put(v)
}
