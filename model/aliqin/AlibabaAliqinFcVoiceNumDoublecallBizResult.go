package aliqin

import (
	"sync"
)

// AlibabaAliqinFcVoiceNumDoublecallBizResult 结构体
type AlibabaAliqinFcVoiceNumDoublecallBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcVoiceNumDoublecallBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcVoiceNumDoublecallBizResult)
	},
}

// GetAlibabaAliqinFcVoiceNumDoublecallBizResult() 从对象池中获取AlibabaAliqinFcVoiceNumDoublecallBizResult
func GetAlibabaAliqinFcVoiceNumDoublecallBizResult() *AlibabaAliqinFcVoiceNumDoublecallBizResult {
	return poolAlibabaAliqinFcVoiceNumDoublecallBizResult.Get().(*AlibabaAliqinFcVoiceNumDoublecallBizResult)
}

// ReleaseAlibabaAliqinFcVoiceNumDoublecallBizResult 释放AlibabaAliqinFcVoiceNumDoublecallBizResult
func ReleaseAlibabaAliqinFcVoiceNumDoublecallBizResult(v *AlibabaAliqinFcVoiceNumDoublecallBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinFcVoiceNumDoublecallBizResult.Put(v)
}
