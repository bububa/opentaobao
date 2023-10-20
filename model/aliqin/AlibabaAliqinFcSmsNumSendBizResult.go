package aliqin

import (
	"sync"
)

// AlibabaAliqinFcSmsNumSendBizResult 结构体
type AlibabaAliqinFcSmsNumSendBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcSmsNumSendBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcSmsNumSendBizResult)
	},
}

// GetAlibabaAliqinFcSmsNumSendBizResult() 从对象池中获取AlibabaAliqinFcSmsNumSendBizResult
func GetAlibabaAliqinFcSmsNumSendBizResult() *AlibabaAliqinFcSmsNumSendBizResult {
	return poolAlibabaAliqinFcSmsNumSendBizResult.Get().(*AlibabaAliqinFcSmsNumSendBizResult)
}

// ReleaseAlibabaAliqinFcSmsNumSendBizResult 释放AlibabaAliqinFcSmsNumSendBizResult
func ReleaseAlibabaAliqinFcSmsNumSendBizResult(v *AlibabaAliqinFcSmsNumSendBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinFcSmsNumSendBizResult.Put(v)
}
