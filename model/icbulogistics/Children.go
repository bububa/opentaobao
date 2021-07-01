package icbulogistics

// Children 结构体
type Children struct {
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品特性列表对象
	Children []LeafNode `json:"children,omitempty" xml:"children>leaf_node,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
