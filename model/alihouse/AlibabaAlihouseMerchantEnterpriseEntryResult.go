package alihouse

import (
	"sync"
)

// AlibabaAlihouseMerchantEnterpriseEntryResult 结构体
type AlibabaAlihouseMerchantEnterpriseEntryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseMerchantEnterpriseEntryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMerchantEnterpriseEntryResult)
	},
}

// GetAlibabaAlihouseMerchantEnterpriseEntryResult() 从对象池中获取AlibabaAlihouseMerchantEnterpriseEntryResult
func GetAlibabaAlihouseMerchantEnterpriseEntryResult() *AlibabaAlihouseMerchantEnterpriseEntryResult {
	return poolAlibabaAlihouseMerchantEnterpriseEntryResult.Get().(*AlibabaAlihouseMerchantEnterpriseEntryResult)
}

// ReleaseAlibabaAlihouseMerchantEnterpriseEntryResult 释放AlibabaAlihouseMerchantEnterpriseEntryResult
func ReleaseAlibabaAlihouseMerchantEnterpriseEntryResult(v *AlibabaAlihouseMerchantEnterpriseEntryResult) {
	v.Message = ""
	v.Code = ""
	v.Data = false
	v.Success = false
	poolAlibabaAlihouseMerchantEnterpriseEntryResult.Put(v)
}
