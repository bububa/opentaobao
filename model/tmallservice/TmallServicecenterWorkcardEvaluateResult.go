package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardEvaluateResult 结构体
type TmallServicecenterWorkcardEvaluateResult struct {
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardEvaluateResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardEvaluateResult)
	},
}

// GetTmallServicecenterWorkcardEvaluateResult() 从对象池中获取TmallServicecenterWorkcardEvaluateResult
func GetTmallServicecenterWorkcardEvaluateResult() *TmallServicecenterWorkcardEvaluateResult {
	return poolTmallServicecenterWorkcardEvaluateResult.Get().(*TmallServicecenterWorkcardEvaluateResult)
}

// ReleaseTmallServicecenterWorkcardEvaluateResult 释放TmallServicecenterWorkcardEvaluateResult
func ReleaseTmallServicecenterWorkcardEvaluateResult(v *TmallServicecenterWorkcardEvaluateResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolTmallServicecenterWorkcardEvaluateResult.Put(v)
}
