package alihouse

import (
	"sync"
)

// AlibabaAlihouseImReceiveModelSyncResult 结构体
type AlibabaAlihouseImReceiveModelSyncResult struct {
	// aaa
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// aaa
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// aaa
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// aaa
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseImReceiveModelSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseImReceiveModelSyncResult)
	},
}

// GetAlibabaAlihouseImReceiveModelSyncResult() 从对象池中获取AlibabaAlihouseImReceiveModelSyncResult
func GetAlibabaAlihouseImReceiveModelSyncResult() *AlibabaAlihouseImReceiveModelSyncResult {
	return poolAlibabaAlihouseImReceiveModelSyncResult.Get().(*AlibabaAlihouseImReceiveModelSyncResult)
}

// ReleaseAlibabaAlihouseImReceiveModelSyncResult 释放AlibabaAlihouseImReceiveModelSyncResult
func ReleaseAlibabaAlihouseImReceiveModelSyncResult(v *AlibabaAlihouseImReceiveModelSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseImReceiveModelSyncResult.Put(v)
}
