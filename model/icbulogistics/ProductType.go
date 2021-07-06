package icbulogistics

// ProductType 结构体
type ProductType struct {
	// 商品特性列表对象
	Children []Children `json:"children,omitempty" xml:"children>children,omitempty"`
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
