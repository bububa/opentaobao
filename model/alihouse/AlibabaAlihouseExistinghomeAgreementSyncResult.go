package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeAgreementSyncResult 结构体
type AlibabaAlihouseExistinghomeAgreementSyncResult struct {
	// 预览URL
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeAgreementSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeAgreementSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeAgreementSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeAgreementSyncResult
func GetAlibabaAlihouseExistinghomeAgreementSyncResult() *AlibabaAlihouseExistinghomeAgreementSyncResult {
	return poolAlibabaAlihouseExistinghomeAgreementSyncResult.Get().(*AlibabaAlihouseExistinghomeAgreementSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeAgreementSyncResult 释放AlibabaAlihouseExistinghomeAgreementSyncResult
func ReleaseAlibabaAlihouseExistinghomeAgreementSyncResult(v *AlibabaAlihouseExistinghomeAgreementSyncResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseExistinghomeAgreementSyncResult.Put(v)
}
