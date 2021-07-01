package trade

// TradeFlowPaymentModel 结构体
type TradeFlowPaymentModel struct {
	// 支付金额
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 支付类型：1扣款，2退款
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 外部系统支付用户ID，比如：支付宝PID
	PaymentUserId string `json:"payment_user_id,omitempty" xml:"payment_user_id,omitempty"`
	// 外部系统支付流水号
	PaymentFlowNo string `json:"payment_flow_no,omitempty" xml:"payment_flow_no,omitempty"`
	// 分佣佣金(单位:分)
	Commission int64 `json:"commission,omitempty" xml:"commission,omitempty"`
	// 支付通道：1现金，2支付宝，3微信，255其他
	PayChannel int64 `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 支付状态：1待付款，2已付款
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 支付创建时间(如果不传将和订单时间相同)
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 支付修改时间(如果不传将和订单时间相同)
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}
