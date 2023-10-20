package scs

import (
	"sync"
)

// DmpJxTopicViewDto 结构体
type DmpJxTopicViewDto struct {
	// template
	TemplateList []DmpJxCrowdTemplateViewDto `json:"template_list,omitempty" xml:"template_list>dmp_jx_crowd_template_view_dto,omitempty"`
	// topicId
	TopicId int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
}

var poolDmpJxTopicViewDto = sync.Pool{
	New: func() any {
		return new(DmpJxTopicViewDto)
	},
}

// GetDmpJxTopicViewDto() 从对象池中获取DmpJxTopicViewDto
func GetDmpJxTopicViewDto() *DmpJxTopicViewDto {
	return poolDmpJxTopicViewDto.Get().(*DmpJxTopicViewDto)
}

// ReleaseDmpJxTopicViewDto 释放DmpJxTopicViewDto
func ReleaseDmpJxTopicViewDto(v *DmpJxTopicViewDto) {
	v.TemplateList = v.TemplateList[:0]
	v.TopicId = 0
	poolDmpJxTopicViewDto.Put(v)
}
