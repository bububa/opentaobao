package idle

import (
	"sync"
)

// AlibabaIdleAgreementPayResult 结构体
type AlibabaIdleAgreementPayResult struct {
	// 错误code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 服务出参
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleAgreementPayResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAgreementPayResult)
	},
}

// GetAlibabaIdleAgreementPayResult() 从对象池中获取AlibabaIdleAgreementPayResult
func GetAlibabaIdleAgreementPayResult() *AlibabaIdleAgreementPayResult {
	return poolAlibabaIdleAgreementPayResult.Get().(*AlibabaIdleAgreementPayResult)
}

// ReleaseAlibabaIdleAgreementPayResult 释放AlibabaIdleAgreementPayResult
func ReleaseAlibabaIdleAgreementPayResult(v *AlibabaIdleAgreementPayResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleAgreementPayResult.Put(v)
}
