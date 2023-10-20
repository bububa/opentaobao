package tmallservice

import (
	"sync"
)

// TmallServicecenterServicestoreUpdateResult 结构体
type TmallServicecenterServicestoreUpdateResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterServicestoreUpdateResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreUpdateResult)
	},
}

// GetTmallServicecenterServicestoreUpdateResult() 从对象池中获取TmallServicecenterServicestoreUpdateResult
func GetTmallServicecenterServicestoreUpdateResult() *TmallServicecenterServicestoreUpdateResult {
	return poolTmallServicecenterServicestoreUpdateResult.Get().(*TmallServicecenterServicestoreUpdateResult)
}

// ReleaseTmallServicecenterServicestoreUpdateResult 释放TmallServicecenterServicestoreUpdateResult
func ReleaseTmallServicecenterServicestoreUpdateResult(v *TmallServicecenterServicestoreUpdateResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTmallServicecenterServicestoreUpdateResult.Put(v)
}
