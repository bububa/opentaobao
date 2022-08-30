package ascp

// DeliveryTemplateModel 结构体
type DeliveryTemplateModel struct {
	// 运费模板名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 运费模板id
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}
