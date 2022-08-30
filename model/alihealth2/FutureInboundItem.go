package alihealth2

// FutureInboundItem 结构体
type FutureInboundItem struct {
	// 期货入库数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
