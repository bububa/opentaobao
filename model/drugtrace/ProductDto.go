package drugtrace

// ProductDto 结构体
type ProductDto struct {
	// 备注
	Comment string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 子类
	SubTypeList []SubType `json:"sub_type_list,omitempty" xml:"sub_type_list>sub_type,omitempty"`
	// 产品 码
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 药品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
}
