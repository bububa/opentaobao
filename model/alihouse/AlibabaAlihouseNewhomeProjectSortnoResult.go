package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectSortnoResult 结构体
type AlibabaAlihouseNewhomeProjectSortnoResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectSortnoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectSortnoResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectSortnoResult() 从对象池中获取AlibabaAlihouseNewhomeProjectSortnoResult
func GetAlibabaAlihouseNewhomeProjectSortnoResult() *AlibabaAlihouseNewhomeProjectSortnoResult {
	return poolAlibabaAlihouseNewhomeProjectSortnoResult.Get().(*AlibabaAlihouseNewhomeProjectSortnoResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectSortnoResult 释放AlibabaAlihouseNewhomeProjectSortnoResult
func ReleaseAlibabaAlihouseNewhomeProjectSortnoResult(v *AlibabaAlihouseNewhomeProjectSortnoResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeProjectSortnoResult.Put(v)
}
