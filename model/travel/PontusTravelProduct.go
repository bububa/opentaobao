package travel

// PontusTravelProduct 结构体
type PontusTravelProduct struct {
	// 资源元素的外部商家编码
	ElementId string `json:"element_id,omitempty" xml:"element_id,omitempty"`
	// 元素份数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
