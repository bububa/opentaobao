package scbp

// ProductReportDTO 
type ProductReportDTO struct {
    // 总数量
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 返回实体集合
    ProductEffectList   []ProductEffectDTO `json:"product_effect_list,omitempty" xml:"product_effect_list>product_effect_dto,omitempty"`
}
