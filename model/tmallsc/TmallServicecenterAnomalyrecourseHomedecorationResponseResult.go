package tmallsc

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationResponseResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationResponseResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationResponseResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationResponseResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationResponseResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationResponseResult
func GetTmallServicecenterAnomalyrecourseHomedecorationResponseResult() *TmallServicecenterAnomalyrecourseHomedecorationResponseResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationResponseResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationResponseResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationResponseResult 释放TmallServicecenterAnomalyrecourseHomedecorationResponseResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationResponseResult(v *TmallServicecenterAnomalyrecourseHomedecorationResponseResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationResponseResult.Put(v)
}
