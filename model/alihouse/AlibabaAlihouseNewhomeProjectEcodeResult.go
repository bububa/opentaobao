package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectEcodeResult 结构体
type AlibabaAlihouseNewhomeProjectEcodeResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectEcodeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectEcodeResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectEcodeResult() 从对象池中获取AlibabaAlihouseNewhomeProjectEcodeResult
func GetAlibabaAlihouseNewhomeProjectEcodeResult() *AlibabaAlihouseNewhomeProjectEcodeResult {
	return poolAlibabaAlihouseNewhomeProjectEcodeResult.Get().(*AlibabaAlihouseNewhomeProjectEcodeResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectEcodeResult 释放AlibabaAlihouseNewhomeProjectEcodeResult
func ReleaseAlibabaAlihouseNewhomeProjectEcodeResult(v *AlibabaAlihouseNewhomeProjectEcodeResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectEcodeResult.Put(v)
}
