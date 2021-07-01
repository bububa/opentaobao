package consignplatform

// OrderCancelRequest 结构体
type OrderCancelRequest struct {
	// 外部订单id
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 外部订单id列表
	SubOuterOrderIds []string `json:"sub_outer_order_ids,omitempty" xml:"sub_outer_order_ids>string,omitempty"`
	// 订单来源
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
}
