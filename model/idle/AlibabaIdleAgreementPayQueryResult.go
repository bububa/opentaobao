package idle

import (
	"sync"
)

// AlibabaIdleAgreementPayQueryResult 结构体
type AlibabaIdleAgreementPayQueryResult struct {
	// err_code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 账单查询结果
	Module *Serializable `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleAgreementPayQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleAgreementPayQueryResult)
	},
}

// GetAlibabaIdleAgreementPayQueryResult() 从对象池中获取AlibabaIdleAgreementPayQueryResult
func GetAlibabaIdleAgreementPayQueryResult() *AlibabaIdleAgreementPayQueryResult {
	return poolAlibabaIdleAgreementPayQueryResult.Get().(*AlibabaIdleAgreementPayQueryResult)
}

// ReleaseAlibabaIdleAgreementPayQueryResult 释放AlibabaIdleAgreementPayQueryResult
func ReleaseAlibabaIdleAgreementPayQueryResult(v *AlibabaIdleAgreementPayQueryResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Module = nil
	v.Success = false
	poolAlibabaIdleAgreementPayQueryResult.Put(v)
}
