package waybill

// UserTemplateDo 
/* model for simplify = false
type UserTemplateDo struct {

    // keys
    
    Keys  struct {
        KeyResult  []KeyResult `json:"key_result,omitempty"`
    } `json:"keys,omitempty"`
    

    // 用户使用模板的url
    
    UserStdTemplateUrl   string `json:"user_std_template_url,omitempty"`
    

    // 用户使用模板的id
    
    UserStdTemplateId   int64 `json:"user_std_template_id,omitempty"`
    

    // 用户使用模板名称
    
    UserStdTemplateName   string `json:"user_std_template_name,omitempty"`
    

}
*/

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
