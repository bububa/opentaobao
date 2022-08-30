package icbulogistics

// SpecialProductTypeDto 结构体
type SpecialProductTypeDto struct {
	// 列表对象
	Childrens []SpecialProductTypeDto `json:"childrens,omitempty" xml:"childrens>special_product_type_dto,omitempty"`
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
