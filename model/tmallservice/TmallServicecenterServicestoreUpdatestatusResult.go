package tmallservice

import (
	"sync"
)

// TmallServicecenterServicestoreUpdatestatusResult 结构体
type TmallServicecenterServicestoreUpdatestatusResult struct {
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterServicestoreUpdatestatusResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreUpdatestatusResult)
	},
}

// GetTmallServicecenterServicestoreUpdatestatusResult() 从对象池中获取TmallServicecenterServicestoreUpdatestatusResult
func GetTmallServicecenterServicestoreUpdatestatusResult() *TmallServicecenterServicestoreUpdatestatusResult {
	return poolTmallServicecenterServicestoreUpdatestatusResult.Get().(*TmallServicecenterServicestoreUpdatestatusResult)
}

// ReleaseTmallServicecenterServicestoreUpdatestatusResult 释放TmallServicecenterServicestoreUpdatestatusResult
func ReleaseTmallServicecenterServicestoreUpdatestatusResult(v *TmallServicecenterServicestoreUpdatestatusResult) {
	v.MsgInfo = ""
	v.MsgCode = 0
	v.Success = false
	poolTmallServicecenterServicestoreUpdatestatusResult.Put(v)
}
