package waybill

// UserTemplateResult 结构体
type UserTemplateResult struct {
	// 用户使用的模板数据
	UserStdTemplates []UserTemplateDo `json:"user_std_templates,omitempty" xml:"user_std_templates>user_template_do,omitempty"`
	// cp编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
}
