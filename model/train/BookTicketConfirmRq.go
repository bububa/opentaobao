package train

// BookTicketConfirmRq 结构体
type BookTicketConfirmRq struct {
	// 票信息
	Tickets []TicketDto `json:"tickets,omitempty" xml:"tickets>ticket_dto,omitempty"`
	// 支付宝交易流水号
	AliPayTradeNo string `json:"ali_pay_trade_no,omitempty" xml:"ali_pay_trade_no,omitempty"`
	// 支付宝账号
	AliPayAccount string `json:"ali_pay_account,omitempty" xml:"ali_pay_account,omitempty"`
	// 检票口信息
	FlatMsg string `json:"flat_msg,omitempty" xml:"flat_msg,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
	// 支付方式 CASH(0),ELECTRONIC(1)
	VipSettlementModeCode int64 `json:"vip_settlement_mode_code,omitempty" xml:"vip_settlement_mode_code,omitempty"`
	// 是否支持线上退改签
	CanChange bool `json:"can_change,omitempty" xml:"can_change,omitempty"`
}
