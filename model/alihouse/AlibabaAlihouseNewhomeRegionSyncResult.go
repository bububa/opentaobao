package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeRegionSyncResult 结构体
type AlibabaAlihouseNewhomeRegionSyncResult struct {
	// 返回外部ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeRegionSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRegionSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeRegionSyncResult() 从对象池中获取AlibabaAlihouseNewhomeRegionSyncResult
func GetAlibabaAlihouseNewhomeRegionSyncResult() *AlibabaAlihouseNewhomeRegionSyncResult {
	return poolAlibabaAlihouseNewhomeRegionSyncResult.Get().(*AlibabaAlihouseNewhomeRegionSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeRegionSyncResult 释放AlibabaAlihouseNewhomeRegionSyncResult
func ReleaseAlibabaAlihouseNewhomeRegionSyncResult(v *AlibabaAlihouseNewhomeRegionSyncResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeRegionSyncResult.Put(v)
}
