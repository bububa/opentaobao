package scs

import (
	"sync"
)

// DmpJxCrowdTemplateViewDto 结构体
type DmpJxCrowdTemplateViewDto struct {
	// 模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 模板有效期
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 模板介绍
	TemplateDesc string `json:"template_desc,omitempty" xml:"template_desc,omitempty"`
	// 模板id:templateId
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

var poolDmpJxCrowdTemplateViewDto = sync.Pool{
	New: func() any {
		return new(DmpJxCrowdTemplateViewDto)
	},
}

// GetDmpJxCrowdTemplateViewDto() 从对象池中获取DmpJxCrowdTemplateViewDto
func GetDmpJxCrowdTemplateViewDto() *DmpJxCrowdTemplateViewDto {
	return poolDmpJxCrowdTemplateViewDto.Get().(*DmpJxCrowdTemplateViewDto)
}

// ReleaseDmpJxCrowdTemplateViewDto 释放DmpJxCrowdTemplateViewDto
func ReleaseDmpJxCrowdTemplateViewDto(v *DmpJxCrowdTemplateViewDto) {
	v.TemplateName = ""
	v.ValidDate = ""
	v.TemplateDesc = ""
	v.TemplateId = 0
	poolDmpJxCrowdTemplateViewDto.Put(v)
}
