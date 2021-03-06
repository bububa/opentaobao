package admarket

// Ad 结构体
type Ad struct {
	// 广告模板id
	TemplateId string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// 物料
	Adm string `json:"adm,omitempty" xml:"adm,omitempty"`
	// 价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 广告目标对象
	Target *Target `json:"target,omitempty" xml:"target,omitempty"`
	// 监控对象
	Monitor *Monitor `json:"monitor,omitempty" xml:"monitor,omitempty"`
}
