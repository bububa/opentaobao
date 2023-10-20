package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardRepairprogressUpdateResult 结构体
type TmallServicecenterWorkcardRepairprogressUpdateResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardRepairprogressUpdateResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardRepairprogressUpdateResult)
	},
}

// GetTmallServicecenterWorkcardRepairprogressUpdateResult() 从对象池中获取TmallServicecenterWorkcardRepairprogressUpdateResult
func GetTmallServicecenterWorkcardRepairprogressUpdateResult() *TmallServicecenterWorkcardRepairprogressUpdateResult {
	return poolTmallServicecenterWorkcardRepairprogressUpdateResult.Get().(*TmallServicecenterWorkcardRepairprogressUpdateResult)
}

// ReleaseTmallServicecenterWorkcardRepairprogressUpdateResult 释放TmallServicecenterWorkcardRepairprogressUpdateResult
func ReleaseTmallServicecenterWorkcardRepairprogressUpdateResult(v *TmallServicecenterWorkcardRepairprogressUpdateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTmallServicecenterWorkcardRepairprogressUpdateResult.Put(v)
}
