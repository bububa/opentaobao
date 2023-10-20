package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeSignatureSyncResult 结构体
type AlibabaAlihouseExistinghomeSignatureSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回素材id
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeSignatureSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeSignatureSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeSignatureSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeSignatureSyncResult
func GetAlibabaAlihouseExistinghomeSignatureSyncResult() *AlibabaAlihouseExistinghomeSignatureSyncResult {
	return poolAlibabaAlihouseExistinghomeSignatureSyncResult.Get().(*AlibabaAlihouseExistinghomeSignatureSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeSignatureSyncResult 释放AlibabaAlihouseExistinghomeSignatureSyncResult
func ReleaseAlibabaAlihouseExistinghomeSignatureSyncResult(v *AlibabaAlihouseExistinghomeSignatureSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihouseExistinghomeSignatureSyncResult.Put(v)
}
