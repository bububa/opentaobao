package alsc

// PaymentOpenInfo 结构体
type PaymentOpenInfo struct {
	// 支付单业务上下文
	BizContext string `json:"biz_context,omitempty" xml:"biz_context,omitempty"`
	// 支付单扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 付款者账户号
	PayerAccountNo string `json:"payer_account_no,omitempty" xml:"payer_account_no,omitempty"`
	// 支付状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 交易类型
	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 支付单资金明细
	FundOpenInfo *FundOpenInfo `json:"fund_open_info,omitempty" xml:"fund_open_info,omitempty"`
}
