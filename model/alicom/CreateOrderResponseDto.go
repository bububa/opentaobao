package alicom

// CreateOrderResponseDto 结构体
type CreateOrderResponseDto struct {
	// 天猫订单号
	TmallOrderId string `json:"tmall_order_id,omitempty" xml:"tmall_order_id,omitempty"`
	// 交易请求流水号
	TransferId string `json:"transfer_id,omitempty" xml:"transfer_id,omitempty"`
	// 金额(单位:分)
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 支付跳转地址
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// 加密串
	SignStr string `json:"sign_str,omitempty" xml:"sign_str,omitempty"`
	// 支付宝支付订单号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
}
