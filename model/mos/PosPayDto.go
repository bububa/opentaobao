package mos

// PosPayDto 结构体
type PosPayDto struct {
	// 订单号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 第三方订单号（如支付宝）
	ThirdTradeNo string `json:"third_trade_no,omitempty" xml:"third_trade_no,omitempty"`
	// 原订单号（退货用）
	OriginalTradeNo string `json:"original_trade_no,omitempty" xml:"original_trade_no,omitempty"`
	// 原外部订单号（退货用）
	OriginalOutTradeNo string `json:"original_out_trade_no,omitempty" xml:"original_out_trade_no,omitempty"`
	// 消费者昵称
	CustomerNickname string `json:"customer_nickname,omitempty" xml:"customer_nickname,omitempty"`
	// 订单金额（支付金额）
	PayAmount int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 支付渠道
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 终端类型（销售渠道）
	SaleChannel int64 `json:"sale_channel,omitempty" xml:"sale_channel,omitempty"`
	// 外部小票号（11位小票号）
	OutTradeNo string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
	// 门店号
	OutBelongId string `json:"out_belong_id,omitempty" xml:"out_belong_id,omitempty"`
	// 专柜号
	OutStoreId string `json:"out_store_id,omitempty" xml:"out_store_id,omitempty"`
	// 收银机号
	TerminalNo string `json:"terminal_no,omitempty" xml:"terminal_no,omitempty"`
	// 收银员
	Cashier string `json:"cashier,omitempty" xml:"cashier,omitempty"`
	// SellerID（如支付宝账号）
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// IP
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// MAC
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// 扩展字段（支付渠道等）
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 交易类型
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
}
