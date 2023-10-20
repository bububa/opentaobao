package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeBankcardSyncResult 结构体
type AlibabaAlihouseExistinghomeBankcardSyncResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 唯一id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeBankcardSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBankcardSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeBankcardSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeBankcardSyncResult
func GetAlibabaAlihouseExistinghomeBankcardSyncResult() *AlibabaAlihouseExistinghomeBankcardSyncResult {
	return poolAlibabaAlihouseExistinghomeBankcardSyncResult.Get().(*AlibabaAlihouseExistinghomeBankcardSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeBankcardSyncResult 释放AlibabaAlihouseExistinghomeBankcardSyncResult
func ReleaseAlibabaAlihouseExistinghomeBankcardSyncResult(v *AlibabaAlihouseExistinghomeBankcardSyncResult) {
	v.Message = ""
	v.Code = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghomeBankcardSyncResult.Put(v)
}
