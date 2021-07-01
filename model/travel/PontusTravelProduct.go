package travel

// PontusTravelProduct 结构体
type PontusTravelProduct struct {
	// 元素的外部商家编码
	ElementId string `json:"element_id,omitempty" xml:"element_id,omitempty"`
	// 份数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
