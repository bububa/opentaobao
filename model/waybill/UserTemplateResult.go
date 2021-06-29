package waybill

// UserTemplateResult 
type UserTemplateResult struct {
    // cp编码
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 用户使用的模板数据
    UserStdTemplates   []UserTemplateDO `json:"user_std_templates,omitempty" xml:"user_std_templates>user_template_do,omitempty"`
}
