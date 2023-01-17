package ascpchannel

// ExtOrderCheckRequest 结构体
type ExtOrderCheckRequest struct {
	// 子订单列表
	OutSubOrders []ExtSubOrderCheckRequest `json:"out_sub_orders,omitempty" xml:"out_sub_orders>ext_sub_order_check_request,omitempty"`
	// 外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 收货人信息
	Receiver *ExtOrderReceiverRequest `json:"receiver,omitempty" xml:"receiver,omitempty"`
}
