package scs

// NewDmpTemplateCrowdTopDto 结构体
type NewDmpTemplateCrowdTopDto struct {
	// 接口查询得到的topicId，需一一对应
	TopicId int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 接口查询得到的templateId，需一一对应
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
