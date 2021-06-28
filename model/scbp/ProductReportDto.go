package scbp

// ProductReportDto 
/* model for simplify = false
type ProductReportDto struct {

    // 总数量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 返回实体集合
    
    ProductEffectList  struct {
        ProductEffectDto  []ProductEffectDto `json:"product_effect_dto,omitempty"`
    } `json:"product_effect_list,omitempty"`
    

}
*/

// ProductReportDto 
type ProductReportDto struct {

    // 总数量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 返回实体集合
    ProductEffectList   []ProductEffectDto `json:"product_effect_list,omitempty"`

}
