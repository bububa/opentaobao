package tmallsc

import (
	"sync"
)

// AlibabaServicecenterWorkcardEvaluateResult 结构体
type AlibabaServicecenterWorkcardEvaluateResult struct {
	// 错误原因文案
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误原因
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaServicecenterWorkcardEvaluateResult = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardEvaluateResult)
	},
}

// GetAlibabaServicecenterWorkcardEvaluateResult() 从对象池中获取AlibabaServicecenterWorkcardEvaluateResult
func GetAlibabaServicecenterWorkcardEvaluateResult() *AlibabaServicecenterWorkcardEvaluateResult {
	return poolAlibabaServicecenterWorkcardEvaluateResult.Get().(*AlibabaServicecenterWorkcardEvaluateResult)
}

// ReleaseAlibabaServicecenterWorkcardEvaluateResult 释放AlibabaServicecenterWorkcardEvaluateResult
func ReleaseAlibabaServicecenterWorkcardEvaluateResult(v *AlibabaServicecenterWorkcardEvaluateResult) {
	v.DisplayMsg = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.IsSuccess = false
	poolAlibabaServicecenterWorkcardEvaluateResult.Put(v)
}
