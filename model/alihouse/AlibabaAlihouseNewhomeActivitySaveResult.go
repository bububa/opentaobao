package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeActivitySaveResult 结构体
type AlibabaAlihouseNewhomeActivitySaveResult struct {
	// 返回编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 调用是否成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 保存是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseNewhomeActivitySaveResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeActivitySaveResult)
	},
}

// GetAlibabaAlihouseNewhomeActivitySaveResult() 从对象池中获取AlibabaAlihouseNewhomeActivitySaveResult
func GetAlibabaAlihouseNewhomeActivitySaveResult() *AlibabaAlihouseNewhomeActivitySaveResult {
	return poolAlibabaAlihouseNewhomeActivitySaveResult.Get().(*AlibabaAlihouseNewhomeActivitySaveResult)
}

// ReleaseAlibabaAlihouseNewhomeActivitySaveResult 释放AlibabaAlihouseNewhomeActivitySaveResult
func ReleaseAlibabaAlihouseNewhomeActivitySaveResult(v *AlibabaAlihouseNewhomeActivitySaveResult) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.Data = false
	poolAlibabaAlihouseNewhomeActivitySaveResult.Put(v)
}
