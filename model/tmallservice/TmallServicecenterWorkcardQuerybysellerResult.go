package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardQuerybysellerResult 结构体
type TmallServicecenterWorkcardQuerybysellerResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 分页数据
	ResultData *PagedResult `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardQuerybysellerResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardQuerybysellerResult)
	},
}

// GetTmallServicecenterWorkcardQuerybysellerResult() 从对象池中获取TmallServicecenterWorkcardQuerybysellerResult
func GetTmallServicecenterWorkcardQuerybysellerResult() *TmallServicecenterWorkcardQuerybysellerResult {
	return poolTmallServicecenterWorkcardQuerybysellerResult.Get().(*TmallServicecenterWorkcardQuerybysellerResult)
}

// ReleaseTmallServicecenterWorkcardQuerybysellerResult 释放TmallServicecenterWorkcardQuerybysellerResult
func ReleaseTmallServicecenterWorkcardQuerybysellerResult(v *TmallServicecenterWorkcardQuerybysellerResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.ResultData = nil
	v.Success = false
	poolTmallServicecenterWorkcardQuerybysellerResult.Put(v)
}
