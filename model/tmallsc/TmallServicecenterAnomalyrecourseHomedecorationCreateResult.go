package tmallsc

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationCreateResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationCreateResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 生成的投诉单id
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationCreateResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationCreateResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationCreateResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationCreateResult
func GetTmallServicecenterAnomalyrecourseHomedecorationCreateResult() *TmallServicecenterAnomalyrecourseHomedecorationCreateResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationCreateResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationCreateResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationCreateResult 释放TmallServicecenterAnomalyrecourseHomedecorationCreateResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationCreateResult(v *TmallServicecenterAnomalyrecourseHomedecorationCreateResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Value = 0
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationCreateResult.Put(v)
}
