package tmallsc

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 问题列表json
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult
func GetTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult() *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult 释放TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult(v *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult) {
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Value = ""
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryResult.Put(v)
}
