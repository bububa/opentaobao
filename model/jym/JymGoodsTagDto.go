package jym

// JymGoodsTagDto 结构体
type JymGoodsTagDto struct {
	// 标签名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标签类型，1-服务 2-卖点
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
