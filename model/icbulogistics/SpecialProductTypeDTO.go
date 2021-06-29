package icbulogistics

// SpecialProductTypeDTO 
type SpecialProductTypeDTO struct {
    // 商品类型code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 商品类型
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 列表对象
    Childrens   []SpecialProductTypeDTO `json:"childrens,omitempty" xml:"childrens>special_product_type_dto,omitempty"`
}
