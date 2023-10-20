package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardSuspendResult 结构体
type TmallServicecenterWorkcardSuspendResult struct {
	// 错误原因描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardSuspendResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardSuspendResult)
	},
}

// GetTmallServicecenterWorkcardSuspendResult() 从对象池中获取TmallServicecenterWorkcardSuspendResult
func GetTmallServicecenterWorkcardSuspendResult() *TmallServicecenterWorkcardSuspendResult {
	return poolTmallServicecenterWorkcardSuspendResult.Get().(*TmallServicecenterWorkcardSuspendResult)
}

// ReleaseTmallServicecenterWorkcardSuspendResult 释放TmallServicecenterWorkcardSuspendResult
func ReleaseTmallServicecenterWorkcardSuspendResult(v *TmallServicecenterWorkcardSuspendResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Success = false
	poolTmallServicecenterWorkcardSuspendResult.Put(v)
}
