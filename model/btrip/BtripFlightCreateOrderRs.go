package btrip

// BtripFlightCreateOrderRs 结构体
type BtripFlightCreateOrderRs struct {
	// 支付宝流水
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 最迟支付时间
	LastPayTime string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// 外部订单ID
	DisOrderId string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// 商旅订单号
	BtripOrderId int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 总支付价格
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}
