package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeLineSyncResult 结构体
type AlibabaAlihouseNewhomeLineSyncResult struct {
	// 外部ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeLineSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLineSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeLineSyncResult() 从对象池中获取AlibabaAlihouseNewhomeLineSyncResult
func GetAlibabaAlihouseNewhomeLineSyncResult() *AlibabaAlihouseNewhomeLineSyncResult {
	return poolAlibabaAlihouseNewhomeLineSyncResult.Get().(*AlibabaAlihouseNewhomeLineSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeLineSyncResult 释放AlibabaAlihouseNewhomeLineSyncResult
func ReleaseAlibabaAlihouseNewhomeLineSyncResult(v *AlibabaAlihouseNewhomeLineSyncResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeLineSyncResult.Put(v)
}
