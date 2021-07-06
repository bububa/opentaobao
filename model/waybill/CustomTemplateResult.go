package waybill

// CustomTemplateResult 结构体
type CustomTemplateResult struct {
	// 模板的keys
	Keys []KeyResult `json:"keys,omitempty" xml:"keys>key_result,omitempty"`
	// isv模板的名称
	IsvTemplateName string `json:"isv_template_name,omitempty" xml:"isv_template_name,omitempty"`
	// isv模板的url
	IsvTemplateUrl string `json:"isv_template_url,omitempty" xml:"isv_template_url,omitempty"`
	// 版本号
	Version string `json:"version,omitempty" xml:"version,omitempty"`
	// isv模板的id
	IsvTemplateId int64 `json:"isv_template_id,omitempty" xml:"isv_template_id,omitempty"`
}
