package scs

// DmpJxTopicViewDto 结构体
type DmpJxTopicViewDto struct {
	// template
	TemplateList []DmpJxCrowdTemplateViewDto `json:"template_list,omitempty" xml:"template_list>dmp_jx_crowd_template_view_dto,omitempty"`
	// topicId
	TopicId int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
}
