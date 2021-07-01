package tmic

// OpenOptionResult 结构体
type OpenOptionResult struct {
	// 问卷记录id，用于区别不同的回答记录，在提交回答的时候作为请求参数
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 是否还有下一题
	HasNextQuestion bool `json:"has_next_question,omitempty" xml:"has_next_question,omitempty"`
	// 业务错误提示
	BizErrInfo string `json:"biz_err_info,omitempty" xml:"biz_err_info,omitempty"`
	// 业务错误编码
	BizErrCode string `json:"biz_err_code,omitempty" xml:"biz_err_code,omitempty"`
	// 某一问题对象
	Question *QuestionBo `json:"question,omitempty" xml:"question,omitempty"`
	// 业务是否调用成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
	// 问卷版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}
