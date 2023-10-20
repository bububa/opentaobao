package alicom

import (
	"sync"
)

// AlibabaAliqinFcVoiceNumCancelcallBizResult 结构体
type AlibabaAliqinFcVoiceNumCancelcallBizResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcVoiceNumCancelcallBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceNumCancelcallBizResult)
	},
}

// GetAlibabaAliqinFcVoiceNumCancelcallBizResult() 从对象池中获取AlibabaAliqinFcVoiceNumCancelcallBizResult
func GetAlibabaAliqinFcVoiceNumCancelcallBizResult() *AlibabaAliqinFcVoiceNumCancelcallBizResult {
	return poolAlibabaAliqinFcVoiceNumCancelcallBizResult.Get().(*AlibabaAliqinFcVoiceNumCancelcallBizResult)
}

// ReleaseAlibabaAliqinFcVoiceNumCancelcallBizResult 释放AlibabaAliqinFcVoiceNumCancelcallBizResult
func ReleaseAlibabaAliqinFcVoiceNumCancelcallBizResult(v *AlibabaAliqinFcVoiceNumCancelcallBizResult) {
	v.Model = ""
	v.ErrCode = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcVoiceNumCancelcallBizResult.Put(v)
}
