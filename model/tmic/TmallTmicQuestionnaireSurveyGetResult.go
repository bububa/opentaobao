package tmic

// TmallTmicQuestionnaireSurveyGetResult 结构体
type TmallTmicQuestionnaireSurveyGetResult struct {
	// 问题列表
	QuestionList []QuestionBo `json:"question_list,omitempty" xml:"question_list>question_bo,omitempty"`
	// 问卷相关的logo地址
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 问卷标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 错误提示
	BizErrorInfo string `json:"biz_error_info,omitempty" xml:"biz_error_info,omitempty"`
	// 问卷描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 错误编码
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 问卷记录id，用于区别不同的回答记录，在提交回答的时候作为请求参数
	RecordId int64 `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 问题数量
	QuestionCount int64 `json:"question_count,omitempty" xml:"question_count,omitempty"`
	// 问卷版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 是否调用成功
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
