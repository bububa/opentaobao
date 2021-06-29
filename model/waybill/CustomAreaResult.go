package waybill

// CustomAreaResult 
type CustomAreaResult struct {
    // customAreaId
    CustomAreaId   int64 `json:"custom_area_id,omitempty" xml:"custom_area_id,omitempty"`
    // customAreaUrl
    CustomAreaUrl   string `json:"custom_area_url,omitempty" xml:"custom_area_url,omitempty"`
    // keys
    Keys   []KeyResult `json:"keys,omitempty" xml:"keys>key_result,omitempty"`
    // 标准模板名称
    StandardTemplateId   int64 `json:"standard_template_id,omitempty" xml:"standard_template_id,omitempty"`
    // 标准模板url
    StandardTemplateUrl   string `json:"standard_template_url,omitempty" xml:"standard_template_url,omitempty"`
    // 自定义区名称
    CustomAreaName   string `json:"custom_area_name,omitempty" xml:"custom_area_name,omitempty"`
    // 用户模板id，等同于mystdtemplates.get中返回的用户模板id
    UserTemplateId   int64 `json:"user_template_id,omitempty" xml:"user_template_id,omitempty"`
}
