package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeVirtualshopSyncResult 结构体
type AlibabaAlihouseExistinghomeVirtualshopSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeVirtualshopSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeVirtualshopSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeVirtualshopSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeVirtualshopSyncResult
func GetAlibabaAlihouseExistinghomeVirtualshopSyncResult() *AlibabaAlihouseExistinghomeVirtualshopSyncResult {
	return poolAlibabaAlihouseExistinghomeVirtualshopSyncResult.Get().(*AlibabaAlihouseExistinghomeVirtualshopSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeVirtualshopSyncResult 释放AlibabaAlihouseExistinghomeVirtualshopSyncResult
func ReleaseAlibabaAlihouseExistinghomeVirtualshopSyncResult(v *AlibabaAlihouseExistinghomeVirtualshopSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghomeVirtualshopSyncResult.Put(v)
}
