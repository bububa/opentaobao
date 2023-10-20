package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeCommunityLineResult 结构体
type AlibabaAlihouseNewhomeCommunityLineResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeCommunityLineResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeCommunityLineResult)
	},
}

// GetAlibabaAlihouseNewhomeCommunityLineResult() 从对象池中获取AlibabaAlihouseNewhomeCommunityLineResult
func GetAlibabaAlihouseNewhomeCommunityLineResult() *AlibabaAlihouseNewhomeCommunityLineResult {
	return poolAlibabaAlihouseNewhomeCommunityLineResult.Get().(*AlibabaAlihouseNewhomeCommunityLineResult)
}

// ReleaseAlibabaAlihouseNewhomeCommunityLineResult 释放AlibabaAlihouseNewhomeCommunityLineResult
func ReleaseAlibabaAlihouseNewhomeCommunityLineResult(v *AlibabaAlihouseNewhomeCommunityLineResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeCommunityLineResult.Put(v)
}
