package tmallsc

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationAppealResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationAppealResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationAppealResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationAppealResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationAppealResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationAppealResult
func GetTmallServicecenterAnomalyrecourseHomedecorationAppealResult() *TmallServicecenterAnomalyrecourseHomedecorationAppealResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationAppealResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationAppealResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationAppealResult 释放TmallServicecenterAnomalyrecourseHomedecorationAppealResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationAppealResult(v *TmallServicecenterAnomalyrecourseHomedecorationAppealResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationAppealResult.Put(v)
}
