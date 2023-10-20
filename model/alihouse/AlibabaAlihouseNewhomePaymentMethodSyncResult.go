package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomePaymentMethodSyncResult 结构体
type AlibabaAlihouseNewhomePaymentMethodSyncResult struct {
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomePaymentMethodSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomePaymentMethodSyncResult)
	},
}

// GetAlibabaAlihouseNewhomePaymentMethodSyncResult() 从对象池中获取AlibabaAlihouseNewhomePaymentMethodSyncResult
func GetAlibabaAlihouseNewhomePaymentMethodSyncResult() *AlibabaAlihouseNewhomePaymentMethodSyncResult {
	return poolAlibabaAlihouseNewhomePaymentMethodSyncResult.Get().(*AlibabaAlihouseNewhomePaymentMethodSyncResult)
}

// ReleaseAlibabaAlihouseNewhomePaymentMethodSyncResult 释放AlibabaAlihouseNewhomePaymentMethodSyncResult
func ReleaseAlibabaAlihouseNewhomePaymentMethodSyncResult(v *AlibabaAlihouseNewhomePaymentMethodSyncResult) {
	v.Code = ""
	v.Message = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseNewhomePaymentMethodSyncResult.Put(v)
}
