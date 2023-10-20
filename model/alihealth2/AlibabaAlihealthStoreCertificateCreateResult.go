package alihealth2

import (
	"sync"
)

// AlibabaAlihealthStoreCertificateCreateResult 结构体
type AlibabaAlihealthStoreCertificateCreateResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
	// 创建审批流是否成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

var poolAlibabaAlihealthStoreCertificateCreateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthStoreCertificateCreateResult)
	},
}

// GetAlibabaAlihealthStoreCertificateCreateResult() 从对象池中获取AlibabaAlihealthStoreCertificateCreateResult
func GetAlibabaAlihealthStoreCertificateCreateResult() *AlibabaAlihealthStoreCertificateCreateResult {
	return poolAlibabaAlihealthStoreCertificateCreateResult.Get().(*AlibabaAlihealthStoreCertificateCreateResult)
}

// ReleaseAlibabaAlihealthStoreCertificateCreateResult 释放AlibabaAlihealthStoreCertificateCreateResult
func ReleaseAlibabaAlihealthStoreCertificateCreateResult(v *AlibabaAlihealthStoreCertificateCreateResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Status = false
	v.Module = false
	poolAlibabaAlihealthStoreCertificateCreateResult.Put(v)
}
