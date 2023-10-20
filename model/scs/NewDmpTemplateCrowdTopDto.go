package scs

import (
	"sync"
)

// NewDmpTemplateCrowdTopDto 结构体
type NewDmpTemplateCrowdTopDto struct {
	// 接口查询得到的topicId，需一一对应
	TopicId int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 接口查询得到的templateId，需一一对应
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

var poolNewDmpTemplateCrowdTopDto = sync.Pool{
	New: func() any {
		return new(NewDmpTemplateCrowdTopDto)
	},
}

// GetNewDmpTemplateCrowdTopDto() 从对象池中获取NewDmpTemplateCrowdTopDto
func GetNewDmpTemplateCrowdTopDto() *NewDmpTemplateCrowdTopDto {
	return poolNewDmpTemplateCrowdTopDto.Get().(*NewDmpTemplateCrowdTopDto)
}

// ReleaseNewDmpTemplateCrowdTopDto 释放NewDmpTemplateCrowdTopDto
func ReleaseNewDmpTemplateCrowdTopDto(v *NewDmpTemplateCrowdTopDto) {
	v.TopicId = 0
	v.TemplateId = 0
	poolNewDmpTemplateCrowdTopDto.Put(v)
}
