package logistic

// ExpressModifyAppointTopRequestDto 结构体
type ExpressModifyAppointTopRequestDto struct {
	// 子交易单号
	SubTradeIds []string `json:"sub_trade_ids,omitempty" xml:"sub_trade_ids>string,omitempty"`
	// 应到达日期
	ScDate string `json:"sc_date,omitempty" xml:"sc_date,omitempty"`
	// 交易号
	TradeId string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// 收货人电话
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 改约日期
	OsDate string `json:"os_date,omitempty" xml:"os_date,omitempty"`
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 外部订单号
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
	// 收货人地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 卖家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
