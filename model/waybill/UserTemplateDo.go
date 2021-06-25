package waybill

// UserTemplateDo 
type UserTemplateDo struct {

    // keys
    Keys   []KeyResult `json:"keys,omitempty"`

    // 用户使用模板的url
    UserStdTemplateUrl   string `json:"user_std_template_url,omitempty"`

    // 用户使用模板的id
    UserStdTemplateId   int64 `json:"user_std_template_id,omitempty"`

    // 用户使用模板名称
    UserStdTemplateName   string `json:"user_std_template_name,omitempty"`

}
