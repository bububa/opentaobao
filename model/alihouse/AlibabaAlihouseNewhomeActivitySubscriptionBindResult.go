package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeActivitySubscriptionBindResult 结构体
type AlibabaAlihouseNewhomeActivitySubscriptionBindResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否保存成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeActivitySubscriptionBindResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivitySubscriptionBindResult)
	},
}

// GetAlibabaAlihouseNewhomeActivitySubscriptionBindResult() 从对象池中获取AlibabaAlihouseNewhomeActivitySubscriptionBindResult
func GetAlibabaAlihouseNewhomeActivitySubscriptionBindResult() *AlibabaAlihouseNewhomeActivitySubscriptionBindResult {
	return poolAlibabaAlihouseNewhomeActivitySubscriptionBindResult.Get().(*AlibabaAlihouseNewhomeActivitySubscriptionBindResult)
}

// ReleaseAlibabaAlihouseNewhomeActivitySubscriptionBindResult 释放AlibabaAlihouseNewhomeActivitySubscriptionBindResult
func ReleaseAlibabaAlihouseNewhomeActivitySubscriptionBindResult(v *AlibabaAlihouseNewhomeActivitySubscriptionBindResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeActivitySubscriptionBindResult.Put(v)
}
