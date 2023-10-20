package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeMetroSyncResult 结构体
type AlibabaAlihouseNewhomeMetroSyncResult struct {
	// 外部ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeMetroSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeMetroSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeMetroSyncResult() 从对象池中获取AlibabaAlihouseNewhomeMetroSyncResult
func GetAlibabaAlihouseNewhomeMetroSyncResult() *AlibabaAlihouseNewhomeMetroSyncResult {
	return poolAlibabaAlihouseNewhomeMetroSyncResult.Get().(*AlibabaAlihouseNewhomeMetroSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeMetroSyncResult 释放AlibabaAlihouseNewhomeMetroSyncResult
func ReleaseAlibabaAlihouseNewhomeMetroSyncResult(v *AlibabaAlihouseNewhomeMetroSyncResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeMetroSyncResult.Put(v)
}
