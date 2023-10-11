package alitripmerchant

// DerbyVoucherCardRedeemDto 结构体
type DerbyVoucherCardRedeemDto struct {
	// 兑换号
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
	// 用户协议
	UserAgreement bool `json:"user_agreement,omitempty" xml:"user_agreement,omitempty"`
	// 兑换协议
	ExchangeAgreement bool `json:"exchange_agreement,omitempty" xml:"exchange_agreement,omitempty"`
}
