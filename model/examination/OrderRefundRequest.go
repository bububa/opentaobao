package examination

// OrderRefundRequest 结构体
type OrderRefundRequest struct {
	// 子单信息
	SubOrderList []SubOrderRequest `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_request,omitempty"`
	// 主订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
