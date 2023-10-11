package idle

// IdleItemApiAutoRechargeDo 结构体
type IdleItemApiAutoRechargeDo struct {
	// add：增加；remove：移除。
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 直充附加信息
	TemplateExtraInfo string `json:"template_extra_info,omitempty" xml:"template_extra_info,omitempty"`
	// 直充模板标识
	TemplateCode string `json:"template_code,omitempty" xml:"template_code,omitempty"`
}
