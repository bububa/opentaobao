package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeCasefieldActivitySubmitResult 结构体
type AlibabaAlihouseNewhomeCasefieldActivitySubmitResult struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回对象
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeCasefieldActivitySubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCasefieldActivitySubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeCasefieldActivitySubmitResult() 从对象池中获取AlibabaAlihouseNewhomeCasefieldActivitySubmitResult
func GetAlibabaAlihouseNewhomeCasefieldActivitySubmitResult() *AlibabaAlihouseNewhomeCasefieldActivitySubmitResult {
	return poolAlibabaAlihouseNewhomeCasefieldActivitySubmitResult.Get().(*AlibabaAlihouseNewhomeCasefieldActivitySubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeCasefieldActivitySubmitResult 释放AlibabaAlihouseNewhomeCasefieldActivitySubmitResult
func ReleaseAlibabaAlihouseNewhomeCasefieldActivitySubmitResult(v *AlibabaAlihouseNewhomeCasefieldActivitySubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeCasefieldActivitySubmitResult.Put(v)
}
