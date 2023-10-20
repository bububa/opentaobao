package tmic

import (
	"sync"
)

// OpenOptionResult 结构体
type OpenOptionResult struct {
	// 业务错误提示
	BizErrInfo string `json:"biz_err_info,omitempty" xml:"biz_err_info,omitempty"`
	// 业务错误编码
	BizErrCode string `json:"biz_err_code,omitempty" xml:"biz_err_code,omitempty"`
	// 问卷记录id，用于区别不同的回答记录，在提交回答的时候作为请求参数
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 某一问题对象
	Question *QuestionBo `json:"question,omitempty" xml:"question,omitempty"`
	// 问卷版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 是否还有下一题
	HasNextQuestion bool `json:"has_next_question,omitempty" xml:"has_next_question,omitempty"`
	// 业务是否调用成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

var poolOpenOptionResult = sync.Pool{
	New: func() any {
		return new(OpenOptionResult)
	},
}

// GetOpenOptionResult() 从对象池中获取OpenOptionResult
func GetOpenOptionResult() *OpenOptionResult {
	return poolOpenOptionResult.Get().(*OpenOptionResult)
}

// ReleaseOpenOptionResult 释放OpenOptionResult
func ReleaseOpenOptionResult(v *OpenOptionResult) {
	v.BizErrInfo = ""
	v.BizErrCode = ""
	v.RecordId = 0
	v.Question = nil
	v.Version = 0
	v.HasNextQuestion = false
	v.BizSuccess = false
	poolOpenOptionResult.Put(v)
}
