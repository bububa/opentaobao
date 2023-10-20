package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardReassignResult 结构体
type TmallServicecenterWorkcardReassignResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 改派后的新的工单id
	ResultData int64 `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardReassignResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardReassignResult)
	},
}

// GetTmallServicecenterWorkcardReassignResult() 从对象池中获取TmallServicecenterWorkcardReassignResult
func GetTmallServicecenterWorkcardReassignResult() *TmallServicecenterWorkcardReassignResult {
	return poolTmallServicecenterWorkcardReassignResult.Get().(*TmallServicecenterWorkcardReassignResult)
}

// ReleaseTmallServicecenterWorkcardReassignResult 释放TmallServicecenterWorkcardReassignResult
func ReleaseTmallServicecenterWorkcardReassignResult(v *TmallServicecenterWorkcardReassignResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResultData = 0
	v.Success = false
	poolTmallServicecenterWorkcardReassignResult.Put(v)
}
