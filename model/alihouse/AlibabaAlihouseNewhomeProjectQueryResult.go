package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectQueryResult 结构体
type AlibabaAlihouseNewhomeProjectQueryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *EbbasItemDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectQueryResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectQueryResult() 从对象池中获取AlibabaAlihouseNewhomeProjectQueryResult
func GetAlibabaAlihouseNewhomeProjectQueryResult() *AlibabaAlihouseNewhomeProjectQueryResult {
	return poolAlibabaAlihouseNewhomeProjectQueryResult.Get().(*AlibabaAlihouseNewhomeProjectQueryResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectQueryResult 释放AlibabaAlihouseNewhomeProjectQueryResult
func ReleaseAlibabaAlihouseNewhomeProjectQueryResult(v *AlibabaAlihouseNewhomeProjectQueryResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectQueryResult.Put(v)
}
