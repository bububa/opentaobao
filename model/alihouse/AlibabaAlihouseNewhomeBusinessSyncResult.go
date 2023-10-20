package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeBusinessSyncResult 结构体
type AlibabaAlihouseNewhomeBusinessSyncResult struct {
	// 外部ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeBusinessSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeBusinessSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeBusinessSyncResult() 从对象池中获取AlibabaAlihouseNewhomeBusinessSyncResult
func GetAlibabaAlihouseNewhomeBusinessSyncResult() *AlibabaAlihouseNewhomeBusinessSyncResult {
	return poolAlibabaAlihouseNewhomeBusinessSyncResult.Get().(*AlibabaAlihouseNewhomeBusinessSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeBusinessSyncResult 释放AlibabaAlihouseNewhomeBusinessSyncResult
func ReleaseAlibabaAlihouseNewhomeBusinessSyncResult(v *AlibabaAlihouseNewhomeBusinessSyncResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeBusinessSyncResult.Put(v)
}
