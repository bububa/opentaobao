package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeLayoutSyncResult 结构体
type AlibabaAlihouseNewhomeLayoutSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回户型ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeLayoutSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLayoutSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeLayoutSyncResult() 从对象池中获取AlibabaAlihouseNewhomeLayoutSyncResult
func GetAlibabaAlihouseNewhomeLayoutSyncResult() *AlibabaAlihouseNewhomeLayoutSyncResult {
	return poolAlibabaAlihouseNewhomeLayoutSyncResult.Get().(*AlibabaAlihouseNewhomeLayoutSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeLayoutSyncResult 释放AlibabaAlihouseNewhomeLayoutSyncResult
func ReleaseAlibabaAlihouseNewhomeLayoutSyncResult(v *AlibabaAlihouseNewhomeLayoutSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeLayoutSyncResult.Put(v)
}
