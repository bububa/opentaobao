package waybill

// StandardTemplateResult 
type StandardTemplateResult struct {
    // cp编码
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 该cp的所有标准模板
    StandardTemplates   []StandardTemplateDO `json:"standard_templates,omitempty" xml:"standard_templates>standard_template_do,omitempty"`
}
