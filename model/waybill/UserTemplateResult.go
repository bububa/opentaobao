package waybill

// UserTemplateResult 
/* model for simplify = false
type UserTemplateResult struct {

    // cp编码
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 用户使用的模板数据
    
    UserStdTemplates  struct {
        UserTemplateDo  []UserTemplateDo `json:"user_template_do,omitempty"`
    } `json:"user_std_templates,omitempty"`
    

}
*/

// UserTemplateResult 
type UserTemplateResult struct {

    // cp编码
    CpCode   string `json:"cp_code,omitempty"`

    // 用户使用的模板数据
    UserStdTemplates   []UserTemplateDo `json:"user_std_templates,omitempty"`

}
