package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeAgreementPreshowResult 结构体
type AlibabaAlihouseNewhomeAgreementPreshowResult struct {
	// 值
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeAgreementPreshowResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeAgreementPreshowResult)
	},
}

// GetAlibabaAlihouseNewhomeAgreementPreshowResult() 从对象池中获取AlibabaAlihouseNewhomeAgreementPreshowResult
func GetAlibabaAlihouseNewhomeAgreementPreshowResult() *AlibabaAlihouseNewhomeAgreementPreshowResult {
	return poolAlibabaAlihouseNewhomeAgreementPreshowResult.Get().(*AlibabaAlihouseNewhomeAgreementPreshowResult)
}

// ReleaseAlibabaAlihouseNewhomeAgreementPreshowResult 释放AlibabaAlihouseNewhomeAgreementPreshowResult
func ReleaseAlibabaAlihouseNewhomeAgreementPreshowResult(v *AlibabaAlihouseNewhomeAgreementPreshowResult) {
	v.Data = ""
	v.Message = ""
	v.Code = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeAgreementPreshowResult.Put(v)
}
