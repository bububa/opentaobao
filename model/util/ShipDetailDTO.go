package util

// ShipDetailDto 结构体
type ShipDetailDto struct {
	// 商品
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 0发货成功 1发货失败 10 核销成功 20 核销失败
	ShipStatus int64 `json:"ship_status,omitempty" xml:"ship_status,omitempty"`
	// 外部订单ID
	OuterTradeNo string `json:"outer_trade_no,omitempty" xml:"outer_trade_no,omitempty"`
}
