package bus

// MerchantBusRefundOrderInfo 结构体
type MerchantBusRefundOrderInfo struct {
	// 退票申请列表
	RefundApplyInfoList []MerchantBusRefundApplyInfo `json:"refund_apply_info_list,omitempty" xml:"refund_apply_info_list>merchant_bus_refund_apply_info,omitempty"`
	// 商家订单号
	AgentOrderId string `json:"agent_order_id,omitempty" xml:"agent_order_id,omitempty"`
	// 支付流水号
	AlipayTradeId string `json:"alipay_trade_id,omitempty" xml:"alipay_trade_id,omitempty"`
	// 淘宝订单号
	TpOrderId string `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 飞猪订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}
