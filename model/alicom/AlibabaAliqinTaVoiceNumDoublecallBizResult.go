package alicom

import (
	"sync"
)

// AlibabaAliqinTaVoiceNumDoublecallBizResult 结构体
type AlibabaAliqinTaVoiceNumDoublecallBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinTaVoiceNumDoublecallBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaVoiceNumDoublecallBizResult)
	},
}

// GetAlibabaAliqinTaVoiceNumDoublecallBizResult() 从对象池中获取AlibabaAliqinTaVoiceNumDoublecallBizResult
func GetAlibabaAliqinTaVoiceNumDoublecallBizResult() *AlibabaAliqinTaVoiceNumDoublecallBizResult {
	return poolAlibabaAliqinTaVoiceNumDoublecallBizResult.Get().(*AlibabaAliqinTaVoiceNumDoublecallBizResult)
}

// ReleaseAlibabaAliqinTaVoiceNumDoublecallBizResult 释放AlibabaAliqinTaVoiceNumDoublecallBizResult
func ReleaseAlibabaAliqinTaVoiceNumDoublecallBizResult(v *AlibabaAliqinTaVoiceNumDoublecallBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinTaVoiceNumDoublecallBizResult.Put(v)
}
