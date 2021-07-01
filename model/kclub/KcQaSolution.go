package kclub

// KcQaSolution 结构体
type KcQaSolution struct {
	// 子知识答案摘要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 子知识答案视角
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 子知识答案额外内容
	ExtraContent string `json:"extra_content,omitempty" xml:"extra_content,omitempty"`
	// 子知识答案视角,逗号分隔
	ContentView string `json:"content_view,omitempty" xml:"content_view,omitempty"`
	// 子知识答案类型
	ContentType int64 `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 子知识答案
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 子知识问题id
	QuestionId int64 `json:"question_id,omitempty" xml:"question_id,omitempty"`
	// 子知识答案编辑时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 子知识答案创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 子知识答案id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 子知识答案纯文本
	PlainText string `json:"plain_text,omitempty" xml:"plain_text,omitempty"`
}
