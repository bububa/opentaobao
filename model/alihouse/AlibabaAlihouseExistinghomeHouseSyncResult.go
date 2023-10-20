package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseSyncResult 结构体
type AlibabaAlihouseExistinghomeHouseSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作结果
	Data *SyncHouseResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseSyncResult
func GetAlibabaAlihouseExistinghomeHouseSyncResult() *AlibabaAlihouseExistinghomeHouseSyncResult {
	return poolAlibabaAlihouseExistinghomeHouseSyncResult.Get().(*AlibabaAlihouseExistinghomeHouseSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseSyncResult 释放AlibabaAlihouseExistinghomeHouseSyncResult
func ReleaseAlibabaAlihouseExistinghomeHouseSyncResult(v *AlibabaAlihouseExistinghomeHouseSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeHouseSyncResult.Put(v)
}
