package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeBrandSyncResult 结构体
type AlibabaAlihouseExistinghomeBrandSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeBrandSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrandSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeBrandSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeBrandSyncResult
func GetAlibabaAlihouseExistinghomeBrandSyncResult() *AlibabaAlihouseExistinghomeBrandSyncResult {
	return poolAlibabaAlihouseExistinghomeBrandSyncResult.Get().(*AlibabaAlihouseExistinghomeBrandSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeBrandSyncResult 释放AlibabaAlihouseExistinghomeBrandSyncResult
func ReleaseAlibabaAlihouseExistinghomeBrandSyncResult(v *AlibabaAlihouseExistinghomeBrandSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghomeBrandSyncResult.Put(v)
}
