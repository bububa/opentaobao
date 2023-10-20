package alicom

import (
	"sync"
)

// AlibabaAliqinFcVoiceRecordGeturlResult 结构体
type AlibabaAliqinFcVoiceRecordGeturlResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcVoiceRecordGeturlResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceRecordGeturlResult)
	},
}

// GetAlibabaAliqinFcVoiceRecordGeturlResult() 从对象池中获取AlibabaAliqinFcVoiceRecordGeturlResult
func GetAlibabaAliqinFcVoiceRecordGeturlResult() *AlibabaAliqinFcVoiceRecordGeturlResult {
	return poolAlibabaAliqinFcVoiceRecordGeturlResult.Get().(*AlibabaAliqinFcVoiceRecordGeturlResult)
}

// ReleaseAlibabaAliqinFcVoiceRecordGeturlResult 释放AlibabaAliqinFcVoiceRecordGeturlResult
func ReleaseAlibabaAliqinFcVoiceRecordGeturlResult(v *AlibabaAliqinFcVoiceRecordGeturlResult) {
	v.Model = ""
	v.Code = ""
	v.Msg = ""
	v.Success = false
	poolAlibabaAliqinFcVoiceRecordGeturlResult.Put(v)
}
