package happytrip

// CarpoolInfo 结构体
type CarpoolInfo struct {
	// 0: 未拼成 1:拼成
	CarpoolFlag int64 `json:"carpool_flag,omitempty" xml:"carpool_flag,omitempty"`
	// 拼车单id
	CarpoolOrderId string `json:"carpool_order_id,omitempty" xml:"carpool_order_id,omitempty"`
	// 乘客订单信息
	PassengerOrders []PassengerOrderInfo `json:"passenger_orders,omitempty" xml:"passenger_orders>passenger_order_info,omitempty"`
}
