package alihouse

import (
	"sync"
)

// AlibabaAlihouseStoreCheckResult 结构体
type AlibabaAlihouseStoreCheckResult struct {
	// 结果列表
	Data []CompanyStoreForCheckDto `json:"data,omitempty" xml:"data>company_store_for_check_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseStoreCheckResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseStoreCheckResult)
	},
}

// GetAlibabaAlihouseStoreCheckResult() 从对象池中获取AlibabaAlihouseStoreCheckResult
func GetAlibabaAlihouseStoreCheckResult() *AlibabaAlihouseStoreCheckResult {
	return poolAlibabaAlihouseStoreCheckResult.Get().(*AlibabaAlihouseStoreCheckResult)
}

// ReleaseAlibabaAlihouseStoreCheckResult 释放AlibabaAlihouseStoreCheckResult
func ReleaseAlibabaAlihouseStoreCheckResult(v *AlibabaAlihouseStoreCheckResult) {
	v.Data = v.Data[:0]
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolAlibabaAlihouseStoreCheckResult.Put(v)
}
