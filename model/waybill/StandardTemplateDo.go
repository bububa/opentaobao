package waybill

// StandardTemplateDo 
type StandardTemplateDo struct {
    // 模板id
    StandardTemplateId   int64 `json:"standard_template_id,omitempty" xml:"standard_template_id,omitempty"`
    // 模板名称
    StandardTemplateName   string `json:"standard_template_name,omitempty" xml:"standard_template_name,omitempty"`
    // 模板url
    StandardTemplateUrl   string `json:"standard_template_url,omitempty" xml:"standard_template_url,omitempty"`
    // 1 快递标准面单 ,2 快递三联面单, 3 快递便携式三联单, 4 快运标准面单, 5 快运三联面单, 6 快递一联单
    StandardWaybillType   int64 `json:"standard_waybill_type,omitempty" xml:"standard_waybill_type,omitempty"`
}
