package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeCommunitySubmitResult 结构体
type AlibabaAlihouseNewhomeCommunitySubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *EbbasCommunitySubmitVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeCommunitySubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCommunitySubmitResult)
	},
}

// GetAlibabaAlihouseNewhomeCommunitySubmitResult() 从对象池中获取AlibabaAlihouseNewhomeCommunitySubmitResult
func GetAlibabaAlihouseNewhomeCommunitySubmitResult() *AlibabaAlihouseNewhomeCommunitySubmitResult {
	return poolAlibabaAlihouseNewhomeCommunitySubmitResult.Get().(*AlibabaAlihouseNewhomeCommunitySubmitResult)
}

// ReleaseAlibabaAlihouseNewhomeCommunitySubmitResult 释放AlibabaAlihouseNewhomeCommunitySubmitResult
func ReleaseAlibabaAlihouseNewhomeCommunitySubmitResult(v *AlibabaAlihouseNewhomeCommunitySubmitResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlihouseNewhomeCommunitySubmitResult.Put(v)
}
