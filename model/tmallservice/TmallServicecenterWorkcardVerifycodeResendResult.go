package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardVerifycodeResendResult 结构体
type TmallServicecenterWorkcardVerifycodeResendResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardVerifycodeResendResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardVerifycodeResendResult)
	},
}

// GetTmallServicecenterWorkcardVerifycodeResendResult() 从对象池中获取TmallServicecenterWorkcardVerifycodeResendResult
func GetTmallServicecenterWorkcardVerifycodeResendResult() *TmallServicecenterWorkcardVerifycodeResendResult {
	return poolTmallServicecenterWorkcardVerifycodeResendResult.Get().(*TmallServicecenterWorkcardVerifycodeResendResult)
}

// ReleaseTmallServicecenterWorkcardVerifycodeResendResult 释放TmallServicecenterWorkcardVerifycodeResendResult
func ReleaseTmallServicecenterWorkcardVerifycodeResendResult(v *TmallServicecenterWorkcardVerifycodeResendResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTmallServicecenterWorkcardVerifycodeResendResult.Put(v)
}
