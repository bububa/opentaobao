package waybill

// StandardTemplateResult 
/* model for simplify = false
type StandardTemplateResult struct {

    // cp编码
    
    CpCode   string `json:"cp_code,omitempty"`
    

    // 该cp的所有标准模板
    
    StandardTemplates  struct {
        StandardTemplateDo  []StandardTemplateDo `json:"standard_template_do,omitempty"`
    } `json:"standard_templates,omitempty"`
    

}
*/

// StandardTemplateResult 
type StandardTemplateResult struct {

    // cp编码
    CpCode   string `json:"cp_code,omitempty"`

    // 该cp的所有标准模板
    StandardTemplates   []StandardTemplateDo `json:"standard_templates,omitempty"`

}
