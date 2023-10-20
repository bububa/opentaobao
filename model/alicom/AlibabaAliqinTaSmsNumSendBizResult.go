package alicom

import (
	"sync"
)

// AlibabaAliqinTaSmsNumSendBizResult 结构体
type AlibabaAliqinTaSmsNumSendBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinTaSmsNumSendBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaSmsNumSendBizResult)
	},
}

// GetAlibabaAliqinTaSmsNumSendBizResult() 从对象池中获取AlibabaAliqinTaSmsNumSendBizResult
func GetAlibabaAliqinTaSmsNumSendBizResult() *AlibabaAliqinTaSmsNumSendBizResult {
	return poolAlibabaAliqinTaSmsNumSendBizResult.Get().(*AlibabaAliqinTaSmsNumSendBizResult)
}

// ReleaseAlibabaAliqinTaSmsNumSendBizResult 释放AlibabaAliqinTaSmsNumSendBizResult
func ReleaseAlibabaAliqinTaSmsNumSendBizResult(v *AlibabaAliqinTaSmsNumSendBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinTaSmsNumSendBizResult.Put(v)
}
