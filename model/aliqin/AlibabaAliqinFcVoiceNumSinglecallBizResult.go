package aliqin

import (
	"sync"
)

// AlibabaAliqinFcVoiceNumSinglecallBizResult 结构体
type AlibabaAliqinFcVoiceNumSinglecallBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcVoiceNumSinglecallBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceNumSinglecallBizResult)
	},
}

// GetAlibabaAliqinFcVoiceNumSinglecallBizResult() 从对象池中获取AlibabaAliqinFcVoiceNumSinglecallBizResult
func GetAlibabaAliqinFcVoiceNumSinglecallBizResult() *AlibabaAliqinFcVoiceNumSinglecallBizResult {
	return poolAlibabaAliqinFcVoiceNumSinglecallBizResult.Get().(*AlibabaAliqinFcVoiceNumSinglecallBizResult)
}

// ReleaseAlibabaAliqinFcVoiceNumSinglecallBizResult 释放AlibabaAliqinFcVoiceNumSinglecallBizResult
func ReleaseAlibabaAliqinFcVoiceNumSinglecallBizResult(v *AlibabaAliqinFcVoiceNumSinglecallBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinFcVoiceNumSinglecallBizResult.Put(v)
}
