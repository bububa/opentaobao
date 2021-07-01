package nlp

// Text 结构体
type Text struct {
	// 业务处理ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 文本内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 文本类型1-评论 2-订单留言 9-其他
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
