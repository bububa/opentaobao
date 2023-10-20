package jst

import (
	"sync"
)

// SmsTemplateRequest 结构体
type SmsTemplateRequest struct {
	// 0--验证码 1--短信通知 2-- 推广短信 3--国际/港澳台消息
	TemplateType string `json:"template_type,omitempty" xml:"template_type,omitempty"`
	// 模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 模板内容
	TemplateContent string `json:"template_content,omitempty" xml:"template_content,omitempty"`
	// 模板CODE
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 描述信息
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 1-- 普通模板  2--数字短信模板
	TemplateClass string `json:"template_class,omitempty" xml:"template_class,omitempty"`
}

var poolSmsTemplateRequest = sync.Pool{
	New: func() any {
		return new(SmsTemplateRequest)
	},
}

// GetSmsTemplateRequest() 从对象池中获取SmsTemplateRequest
func GetSmsTemplateRequest() *SmsTemplateRequest {
	return poolSmsTemplateRequest.Get().(*SmsTemplateRequest)
}

// ReleaseSmsTemplateRequest 释放SmsTemplateRequest
func ReleaseSmsTemplateRequest(v *SmsTemplateRequest) {
	v.TemplateType = ""
	v.TemplateName = ""
	v.TemplateContent = ""
	v.TemplateCode = ""
	v.Desc = ""
	v.TemplateClass = ""
	poolSmsTemplateRequest.Put(v)
}
