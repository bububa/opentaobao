package dmp

import (
	"sync"
)

// TemplateContextDto 结构体
type TemplateContextDto struct {
	// 算法推荐模版context
	ContextParams string `json:"context_params,omitempty" xml:"context_params,omitempty"`
	// 榜单id
	TopicId int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 模版id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

var poolTemplateContextDto = sync.Pool{
	New: func() any {
		return new(TemplateContextDto)
	},
}

// GetTemplateContextDto() 从对象池中获取TemplateContextDto
func GetTemplateContextDto() *TemplateContextDto {
	return poolTemplateContextDto.Get().(*TemplateContextDto)
}

// ReleaseTemplateContextDto 释放TemplateContextDto
func ReleaseTemplateContextDto(v *TemplateContextDto) {
	v.ContextParams = ""
	v.TopicId = 0
	v.TemplateId = 0
	poolTemplateContextDto.Put(v)
}
