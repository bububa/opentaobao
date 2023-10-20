package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult 结构体
type AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult)
	},
}

// GetAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult() 从对象池中获取AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult
func GetAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult() *AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult {
	return poolAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult.Get().(*AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult)
}

// ReleaseAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult 释放AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult
func ReleaseAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult(v *AlibabaAlihouseExistinghomeCommunityBrokerSubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseExistinghomeCommunityBrokerSubmitResult.Put(v)
}
