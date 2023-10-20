package jst

import (
	"sync"
)

// AccessBaseDto 结构体
type AccessBaseDto struct {
	// 审核意见
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 模板内容
	TemplateContent string `json:"template_content,omitempty" xml:"template_content,omitempty"`
	// 模板CODE
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 创建时间
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// 0--验证码 1--短信通知 2-- 推广短信 3--国际/港澳台消息
	TemplateType int64 `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// 0--待审核  1--通过  2--拒绝
	TemplateStatus int64 `json:"template_status,omitempty" xml:"template_status,omitempty"`
}

var poolAccessBaseDto = sync.Pool{
	New: func() any {
		return new(AccessBaseDto)
	},
}

// GetAccessBaseDto() 从对象池中获取AccessBaseDto
func GetAccessBaseDto() *AccessBaseDto {
	return poolAccessBaseDto.Get().(*AccessBaseDto)
}

// ReleaseAccessBaseDto 释放AccessBaseDto
func ReleaseAccessBaseDto(v *AccessBaseDto) {
	v.Reason = ""
	v.TemplateName = ""
	v.TemplateContent = ""
	v.TemplateCode = ""
	v.CreateDate = ""
	v.TemplateType = 0
	v.TemplateStatus = 0
	poolAccessBaseDto.Put(v)
}
