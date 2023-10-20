package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardReserveResult 结构体
type TmallServicecenterWorkcardReserveResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardReserveResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardReserveResult)
	},
}

// GetTmallServicecenterWorkcardReserveResult() 从对象池中获取TmallServicecenterWorkcardReserveResult
func GetTmallServicecenterWorkcardReserveResult() *TmallServicecenterWorkcardReserveResult {
	return poolTmallServicecenterWorkcardReserveResult.Get().(*TmallServicecenterWorkcardReserveResult)
}

// ReleaseTmallServicecenterWorkcardReserveResult 释放TmallServicecenterWorkcardReserveResult
func ReleaseTmallServicecenterWorkcardReserveResult(v *TmallServicecenterWorkcardReserveResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTmallServicecenterWorkcardReserveResult.Put(v)
}
