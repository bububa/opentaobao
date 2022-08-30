package jst

// TopModifySmsTemplateRequest 结构体
type TopModifySmsTemplateRequest struct {
	// 使用场景说明，可以修改
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 不能修改
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
	// 模板名称，可以修改
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 模板内容，占位符用${}标识，可以修改
	TemplateContent string `json:"template_content,omitempty" xml:"template_content,omitempty"`
	// 不能修改
	TemplateType int64 `json:"template_type,omitempty" xml:"template_type,omitempty"`
}
