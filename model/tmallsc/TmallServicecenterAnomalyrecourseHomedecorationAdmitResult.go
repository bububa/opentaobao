package tmallsc

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationAdmitResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationAdmitResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationAdmitResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationAdmitResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationAdmitResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationAdmitResult
func GetTmallServicecenterAnomalyrecourseHomedecorationAdmitResult() *TmallServicecenterAnomalyrecourseHomedecorationAdmitResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationAdmitResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationAdmitResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationAdmitResult 释放TmallServicecenterAnomalyrecourseHomedecorationAdmitResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationAdmitResult(v *TmallServicecenterAnomalyrecourseHomedecorationAdmitResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationAdmitResult.Put(v)
}
