package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeSupportSyncResult 结构体
type AlibabaAlihouseNewhomeSupportSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeSupportSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeSupportSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeSupportSyncResult() 从对象池中获取AlibabaAlihouseNewhomeSupportSyncResult
func GetAlibabaAlihouseNewhomeSupportSyncResult() *AlibabaAlihouseNewhomeSupportSyncResult {
	return poolAlibabaAlihouseNewhomeSupportSyncResult.Get().(*AlibabaAlihouseNewhomeSupportSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeSupportSyncResult 释放AlibabaAlihouseNewhomeSupportSyncResult
func ReleaseAlibabaAlihouseNewhomeSupportSyncResult(v *AlibabaAlihouseNewhomeSupportSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseNewhomeSupportSyncResult.Put(v)
}
