package tbk

// OrderData 结构体
type OrderData struct {
	// 预估佣金
	Commission string `json:"commission,omitempty" xml:"commission,omitempty"`
	// 收货时间
	Confirmreceivetime string `json:"confirm_receive_time,omitempty" xml:"confirm_receive_time,omitempty"`
	// 支付时间
	Paytime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 订单号
	Orderno string `json:"order_no,omitempty" xml:"order_no,omitempty"`
}
