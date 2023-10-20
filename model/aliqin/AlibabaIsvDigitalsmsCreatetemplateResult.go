package aliqin

import (
	"sync"
)

// AlibabaIsvDigitalsmsCreatetemplateResult 结构体
type AlibabaIsvDigitalsmsCreatetemplateResult struct {
	// 模板code
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIsvDigitalsmsCreatetemplateResult = sync.Pool{
	New: func() any {
		return new(AlibabaIsvDigitalsmsCreatetemplateResult)
	},
}

// GetAlibabaIsvDigitalsmsCreatetemplateResult() 从对象池中获取AlibabaIsvDigitalsmsCreatetemplateResult
func GetAlibabaIsvDigitalsmsCreatetemplateResult() *AlibabaIsvDigitalsmsCreatetemplateResult {
	return poolAlibabaIsvDigitalsmsCreatetemplateResult.Get().(*AlibabaIsvDigitalsmsCreatetemplateResult)
}

// ReleaseAlibabaIsvDigitalsmsCreatetemplateResult 释放AlibabaIsvDigitalsmsCreatetemplateResult
func ReleaseAlibabaIsvDigitalsmsCreatetemplateResult(v *AlibabaIsvDigitalsmsCreatetemplateResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaIsvDigitalsmsCreatetemplateResult.Put(v)
}
