package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeQuotationSyncResult 结构体
type AlibabaAlihouseExistinghomeQuotationSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否同步成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeQuotationSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeQuotationSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeQuotationSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeQuotationSyncResult
func GetAlibabaAlihouseExistinghomeQuotationSyncResult() *AlibabaAlihouseExistinghomeQuotationSyncResult {
	return poolAlibabaAlihouseExistinghomeQuotationSyncResult.Get().(*AlibabaAlihouseExistinghomeQuotationSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeQuotationSyncResult 释放AlibabaAlihouseExistinghomeQuotationSyncResult
func ReleaseAlibabaAlihouseExistinghomeQuotationSyncResult(v *AlibabaAlihouseExistinghomeQuotationSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseExistinghomeQuotationSyncResult.Put(v)
}
