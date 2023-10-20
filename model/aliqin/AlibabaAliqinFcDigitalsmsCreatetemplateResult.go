package aliqin

import (
	"sync"
)

// AlibabaAliqinFcDigitalsmsCreatetemplateResult 结构体
type AlibabaAliqinFcDigitalsmsCreatetemplateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 模板code
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcDigitalsmsCreatetemplateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcDigitalsmsCreatetemplateResult)
	},
}

// GetAlibabaAliqinFcDigitalsmsCreatetemplateResult() 从对象池中获取AlibabaAliqinFcDigitalsmsCreatetemplateResult
func GetAlibabaAliqinFcDigitalsmsCreatetemplateResult() *AlibabaAliqinFcDigitalsmsCreatetemplateResult {
	return poolAlibabaAliqinFcDigitalsmsCreatetemplateResult.Get().(*AlibabaAliqinFcDigitalsmsCreatetemplateResult)
}

// ReleaseAlibabaAliqinFcDigitalsmsCreatetemplateResult 释放AlibabaAliqinFcDigitalsmsCreatetemplateResult
func ReleaseAlibabaAliqinFcDigitalsmsCreatetemplateResult(v *AlibabaAliqinFcDigitalsmsCreatetemplateResult) {
	v.ErrCode = ""
	v.Msg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaAliqinFcDigitalsmsCreatetemplateResult.Put(v)
}
