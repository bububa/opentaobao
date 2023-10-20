package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeStoreSyncResult 结构体
type AlibabaAlihouseExistinghomeStoreSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeStoreSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeStoreSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeStoreSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeStoreSyncResult
func GetAlibabaAlihouseExistinghomeStoreSyncResult() *AlibabaAlihouseExistinghomeStoreSyncResult {
	return poolAlibabaAlihouseExistinghomeStoreSyncResult.Get().(*AlibabaAlihouseExistinghomeStoreSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeStoreSyncResult 释放AlibabaAlihouseExistinghomeStoreSyncResult
func ReleaseAlibabaAlihouseExistinghomeStoreSyncResult(v *AlibabaAlihouseExistinghomeStoreSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghomeStoreSyncResult.Put(v)
}
