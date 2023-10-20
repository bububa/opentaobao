package tmallsc

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationCloseResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationCloseResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationCloseResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationCloseResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationCloseResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationCloseResult
func GetTmallServicecenterAnomalyrecourseHomedecorationCloseResult() *TmallServicecenterAnomalyrecourseHomedecorationCloseResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationCloseResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationCloseResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationCloseResult 释放TmallServicecenterAnomalyrecourseHomedecorationCloseResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationCloseResult(v *TmallServicecenterAnomalyrecourseHomedecorationCloseResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationCloseResult.Put(v)
}
