package waybill

// StandardTemplateResult 结构体
type StandardTemplateResult struct {
	// 该cp的所有标准模板
	StandardTemplates []StandardTemplateDo `json:"standard_templates,omitempty" xml:"standard_templates>standard_template_do,omitempty"`
	// cp编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}
