package product

// TmallItemSizemappingTemplateGetModel 结构体
type TmallItemSizemappingTemplateGetModel struct {
	// 尺码表模板内容，格式为"尺码值:维度名称:数值,尺码值:维度名称:数值"。其中，数值的单位，长度单位为厘米（cm），体重单位为公斤（kg）。
	TemplateContent string `json:"template_content,omitempty" xml:"template_content,omitempty"`
	// 尺码表模板名称
	TemplateName string `json:"template_name,omitempty" xml:"template_name,omitempty"`
	// 尺码表模板ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
