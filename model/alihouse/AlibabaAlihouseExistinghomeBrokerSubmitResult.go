package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeBrokerSubmitResult 结构体
type AlibabaAlihouseExistinghomeBrokerSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeBrokerSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrokerSubmitResult)
	},
}

// GetAlibabaAlihouseExistinghomeBrokerSubmitResult() 从对象池中获取AlibabaAlihouseExistinghomeBrokerSubmitResult
func GetAlibabaAlihouseExistinghomeBrokerSubmitResult() *AlibabaAlihouseExistinghomeBrokerSubmitResult {
	return poolAlibabaAlihouseExistinghomeBrokerSubmitResult.Get().(*AlibabaAlihouseExistinghomeBrokerSubmitResult)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerSubmitResult 释放AlibabaAlihouseExistinghomeBrokerSubmitResult
func ReleaseAlibabaAlihouseExistinghomeBrokerSubmitResult(v *AlibabaAlihouseExistinghomeBrokerSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseExistinghomeBrokerSubmitResult.Put(v)
}
