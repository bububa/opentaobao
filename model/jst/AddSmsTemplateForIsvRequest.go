package jst

import (
	"sync"
)

// AddSmsTemplateForIsvRequest 结构体
type AddSmsTemplateForIsvRequest struct {
	// 上传文件
	TemplateInfos []DigitalSmsTemplateContentDto `json:"template_infos,omitempty" xml:"template_infos>digital_sms_template_content_dto,omitempty"`
	// NORMAL -- 普通模板  DIGITAL--数字短信模板
	TemplateTypeClass string `json:"template_type_class,omitempty" xml:"template_type_class,omitempty"`
	// 使用场景说明
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 模板内容，占位符用${}标识
	TemplateContent string `json:"template_content,omitempty" xml:"template_content,omitempty"`
	// 0--验证码 1--短信通知 2-- 推广短信 3--国际/港澳台消息
	TemplateType int64 `json:"template_type,omitempty" xml:"template_type,omitempty"`
}

var poolAddSmsTemplateForIsvRequest = sync.Pool{
	New: func() any {
		return new(AddSmsTemplateForIsvRequest)
	},
}

// GetAddSmsTemplateForIsvRequest() 从对象池中获取AddSmsTemplateForIsvRequest
func GetAddSmsTemplateForIsvRequest() *AddSmsTemplateForIsvRequest {
	return poolAddSmsTemplateForIsvRequest.Get().(*AddSmsTemplateForIsvRequest)
}

// ReleaseAddSmsTemplateForIsvRequest 释放AddSmsTemplateForIsvRequest
func ReleaseAddSmsTemplateForIsvRequest(v *AddSmsTemplateForIsvRequest) {
	v.TemplateInfos = v.TemplateInfos[:0]
	v.TemplateTypeClass = ""
	v.Remark = ""
	v.TemplateName = ""
	v.TemplateContent = ""
	v.TemplateType = 0
	poolAddSmsTemplateForIsvRequest.Put(v)
}
