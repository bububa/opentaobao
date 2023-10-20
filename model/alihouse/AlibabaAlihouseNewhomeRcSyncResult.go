package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeRcSyncResult 结构体
type AlibabaAlihouseNewhomeRcSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseNewhomeRcSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeRcSyncResult)
	},
}

// GetAlibabaAlihouseNewhomeRcSyncResult() 从对象池中获取AlibabaAlihouseNewhomeRcSyncResult
func GetAlibabaAlihouseNewhomeRcSyncResult() *AlibabaAlihouseNewhomeRcSyncResult {
	return poolAlibabaAlihouseNewhomeRcSyncResult.Get().(*AlibabaAlihouseNewhomeRcSyncResult)
}

// ReleaseAlibabaAlihouseNewhomeRcSyncResult 释放AlibabaAlihouseNewhomeRcSyncResult
func ReleaseAlibabaAlihouseNewhomeRcSyncResult(v *AlibabaAlihouseNewhomeRcSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseNewhomeRcSyncResult.Put(v)
}
