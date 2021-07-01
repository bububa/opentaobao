package scbp

// ProductReportDto 结构体
type ProductReportDto struct {
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回实体集合
	ProductEffectList []ProductEffectDto `json:"product_effect_list,omitempty" xml:"product_effect_list>product_effect_dto,omitempty"`
}
