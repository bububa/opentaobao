package alihouse

import (
	"sync"
)

// AlibabaAlihouseProjectActivitySyncResult 结构体
type AlibabaAlihouseProjectActivitySyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 主体ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseProjectActivitySyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseProjectActivitySyncResult)
	},
}

// GetAlibabaAlihouseProjectActivitySyncResult() 从对象池中获取AlibabaAlihouseProjectActivitySyncResult
func GetAlibabaAlihouseProjectActivitySyncResult() *AlibabaAlihouseProjectActivitySyncResult {
	return poolAlibabaAlihouseProjectActivitySyncResult.Get().(*AlibabaAlihouseProjectActivitySyncResult)
}

// ReleaseAlibabaAlihouseProjectActivitySyncResult 释放AlibabaAlihouseProjectActivitySyncResult
func ReleaseAlibabaAlihouseProjectActivitySyncResult(v *AlibabaAlihouseProjectActivitySyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseProjectActivitySyncResult.Put(v)
}
