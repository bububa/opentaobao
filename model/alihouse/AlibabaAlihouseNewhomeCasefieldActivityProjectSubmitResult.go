package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult 结构体
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult struct {
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回值
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult() 从对象池中获取AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult
func GetAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult() *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult {
	return poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult.Get().(*AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult 释放AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult
func ReleaseAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult(v *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitResult.Put(v)
}
