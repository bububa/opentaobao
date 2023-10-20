package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeProjectSalestimeResult 结构体
type AlibabaAlihouseNewhomeProjectSalestimeResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeProjectSalestimeResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectSalestimeResult)
	},
}

// GetAlibabaAlihouseNewhomeProjectSalestimeResult() 从对象池中获取AlibabaAlihouseNewhomeProjectSalestimeResult
func GetAlibabaAlihouseNewhomeProjectSalestimeResult() *AlibabaAlihouseNewhomeProjectSalestimeResult {
	return poolAlibabaAlihouseNewhomeProjectSalestimeResult.Get().(*AlibabaAlihouseNewhomeProjectSalestimeResult)
}

// ReleaseAlibabaAlihouseNewhomeProjectSalestimeResult 释放AlibabaAlihouseNewhomeProjectSalestimeResult
func ReleaseAlibabaAlihouseNewhomeProjectSalestimeResult(v *AlibabaAlihouseNewhomeProjectSalestimeResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeProjectSalestimeResult.Put(v)
}
