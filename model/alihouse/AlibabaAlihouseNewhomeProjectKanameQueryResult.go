package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectKanameQueryResult 结构体
type AlibabaAlihouseNewhomeProjectKanameQueryResult struct {
	// 返回对象
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectKanameQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectKanameQueryResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectKanameQueryResult() 从对象池中获取AlibabaAlihouseNewhomeProjectKanameQueryResult
func GetAlibabaAlihouseNewhomeProjectKanameQueryResult() *AlibabaAlihouseNewhomeProjectKanameQueryResult {
	return poolAlibabaAlihouseNewhomeProjectKanameQueryResult.Get().(*AlibabaAlihouseNewhomeProjectKanameQueryResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectKanameQueryResult 释放AlibabaAlihouseNewhomeProjectKanameQueryResult
func ReleaseAlibabaAlihouseNewhomeProjectKanameQueryResult(v *AlibabaAlihouseNewhomeProjectKanameQueryResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectKanameQueryResult.Put(v)
}
