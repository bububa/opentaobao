package aliqin

import (
	"sync"
)

// AlibabaAliyunindepDigitalsmsCreatetemplateResult 结构体
type AlibabaAliyunindepDigitalsmsCreatetemplateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 模板code
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliyunindepDigitalsmsCreatetemplateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliyunindepDigitalsmsCreatetemplateResult)
	},
}

// GetAlibabaAliyunindepDigitalsmsCreatetemplateResult() 从对象池中获取AlibabaAliyunindepDigitalsmsCreatetemplateResult
func GetAlibabaAliyunindepDigitalsmsCreatetemplateResult() *AlibabaAliyunindepDigitalsmsCreatetemplateResult {
	return poolAlibabaAliyunindepDigitalsmsCreatetemplateResult.Get().(*AlibabaAliyunindepDigitalsmsCreatetemplateResult)
}

// ReleaseAlibabaAliyunindepDigitalsmsCreatetemplateResult 释放AlibabaAliyunindepDigitalsmsCreatetemplateResult
func ReleaseAlibabaAliyunindepDigitalsmsCreatetemplateResult(v *AlibabaAliyunindepDigitalsmsCreatetemplateResult) {
	v.ErrCode = ""
	v.Msg = ""
	v.Model = ""
	v.Success = false
	poolAlibabaAliyunindepDigitalsmsCreatetemplateResult.Put(v)
}
