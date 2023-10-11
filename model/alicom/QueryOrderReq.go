package alicom

// QueryOrderReq 结构体
type QueryOrderReq struct {
	// 外部订单id
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}
