package aliqin

import (
	"sync"
)

// AlibabaAliqinFcIotSmsSendBizResult 结构体
type AlibabaAliqinFcIotSmsSendBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcIotSmsSendBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotSmsSendBizResult)
	},
}

// GetAlibabaAliqinFcIotSmsSendBizResult() 从对象池中获取AlibabaAliqinFcIotSmsSendBizResult
func GetAlibabaAliqinFcIotSmsSendBizResult() *AlibabaAliqinFcIotSmsSendBizResult {
	return poolAlibabaAliqinFcIotSmsSendBizResult.Get().(*AlibabaAliqinFcIotSmsSendBizResult)
}

// ReleaseAlibabaAliqinFcIotSmsSendBizResult 释放AlibabaAliqinFcIotSmsSendBizResult
func ReleaseAlibabaAliqinFcIotSmsSendBizResult(v *AlibabaAliqinFcIotSmsSendBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinFcIotSmsSendBizResult.Put(v)
}
