package waybill

// CustomTemplateResult 
type CustomTemplateResult struct {

    // isv模板的id
    IsvTemplateId   int64 `json:"isv_template_id,omitempty"`

    // isv模板的名称
    IsvTemplateName   string `json:"isv_template_name,omitempty"`

    // isv模板的url
    IsvTemplateUrl   string `json:"isv_template_url,omitempty"`

    // 模板的keys
    Keys   []KeyResult `json:"keys,omitempty"`

    // 版本号
    Version   string `json:"version,omitempty"`

}
