package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardAssignworkerResult 结构体
type TmallServicecenterWorkcardAssignworkerResult struct {
	// 用于对外展示的信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 消息代码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 系统级错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardAssignworkerResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardAssignworkerResult)
	},
}

// GetTmallServicecenterWorkcardAssignworkerResult() 从对象池中获取TmallServicecenterWorkcardAssignworkerResult
func GetTmallServicecenterWorkcardAssignworkerResult() *TmallServicecenterWorkcardAssignworkerResult {
	return poolTmallServicecenterWorkcardAssignworkerResult.Get().(*TmallServicecenterWorkcardAssignworkerResult)
}

// ReleaseTmallServicecenterWorkcardAssignworkerResult 释放TmallServicecenterWorkcardAssignworkerResult
func ReleaseTmallServicecenterWorkcardAssignworkerResult(v *TmallServicecenterWorkcardAssignworkerResult) {
	v.DisplayMsg = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTmallServicecenterWorkcardAssignworkerResult.Put(v)
}
