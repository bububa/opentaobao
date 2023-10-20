package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectAdviserDeleteResult 结构体
type AlibabaAlihouseNewhomeProjectAdviserDeleteResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回记过
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectAdviserDeleteResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectAdviserDeleteResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserDeleteResult() 从对象池中获取AlibabaAlihouseNewhomeProjectAdviserDeleteResult
func GetAlibabaAlihouseNewhomeProjectAdviserDeleteResult() *AlibabaAlihouseNewhomeProjectAdviserDeleteResult {
	return poolAlibabaAlihouseNewhomeProjectAdviserDeleteResult.Get().(*AlibabaAlihouseNewhomeProjectAdviserDeleteResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserDeleteResult 释放AlibabaAlihouseNewhomeProjectAdviserDeleteResult
func ReleaseAlibabaAlihouseNewhomeProjectAdviserDeleteResult(v *AlibabaAlihouseNewhomeProjectAdviserDeleteResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectAdviserDeleteResult.Put(v)
}
