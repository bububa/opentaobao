package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardQueryResult 结构体
type TmallServicecenterWorkcardQueryResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 分页数据
	ResultData *Paged `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardQueryResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardQueryResult)
	},
}

// GetTmallServicecenterWorkcardQueryResult() 从对象池中获取TmallServicecenterWorkcardQueryResult
func GetTmallServicecenterWorkcardQueryResult() *TmallServicecenterWorkcardQueryResult {
	return poolTmallServicecenterWorkcardQueryResult.Get().(*TmallServicecenterWorkcardQueryResult)
}

// ReleaseTmallServicecenterWorkcardQueryResult 释放TmallServicecenterWorkcardQueryResult
func ReleaseTmallServicecenterWorkcardQueryResult(v *TmallServicecenterWorkcardQueryResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResultData = nil
	v.Success = false
	poolTmallServicecenterWorkcardQueryResult.Put(v)
}
