package ascpchannel

// Item 结构体
type Item struct {
	// 后端货品 ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 商家编码字段
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
