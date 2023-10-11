package tmallhk

// SupplierBreakOrderRequest 结构体
type SupplierBreakOrderRequest struct {
	// 毁单商品详细信息
	BrokenOrderItemInfos []BrokenOrderItemInfo `json:"broken_order_item_infos,omitempty" xml:"broken_order_item_infos>broken_order_item_info,omitempty"`
	// 订单毁单时间
	BreakOrderTime string `json:"break_order_time,omitempty" xml:"break_order_time,omitempty"`
	// 毁单操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 主订单信息
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
