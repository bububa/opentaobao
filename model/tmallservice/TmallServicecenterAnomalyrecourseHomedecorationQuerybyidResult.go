package tmallservice

import (
	"sync"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult 结构体
type TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult struct {
	// 投诉单对象的json字符串
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult)
	},
}

// GetTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult() 从对象池中获取TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult
func GetTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult() *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult {
	return poolTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult.Get().(*TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult)
}

// ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult 释放TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult
func ReleaseTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult(v *TmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult) {
	v.Value = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTmallServicecenterAnomalyrecourseHomedecorationQuerybyidResult.Put(v)
}
