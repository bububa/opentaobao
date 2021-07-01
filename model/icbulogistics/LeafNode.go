package icbulogistics

// LeafNode 结构体
type LeafNode struct {
	// 商品类型code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 商品类型
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
