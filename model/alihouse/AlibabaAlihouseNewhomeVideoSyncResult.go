package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeVideoSyncResult 结构体
type AlibabaAlihouseNewhomeVideoSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseNewhomeVideoSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeVideoSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeVideoSyncResult() 从对象池中获取AlibabaAlihouseNewhomeVideoSyncResult
func GetAlibabaAlihouseNewhomeVideoSyncResult() *AlibabaAlihouseNewhomeVideoSyncResult {
	return poolAlibabaAlihouseNewhomeVideoSyncResult.Get().(*AlibabaAlihouseNewhomeVideoSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeVideoSyncResult 释放AlibabaAlihouseNewhomeVideoSyncResult
func ReleaseAlibabaAlihouseNewhomeVideoSyncResult(v *AlibabaAlihouseNewhomeVideoSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseNewhomeVideoSyncResult.Put(v)
}
