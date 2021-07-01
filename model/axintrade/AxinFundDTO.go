package axintrade

// AxinFundDto 结构体
type AxinFundDto struct {
	// 有效资金单列表
	FundList []AxinFundDto `json:"fund_list,omitempty" xml:"fund_list>axin_fund_dto,omitempty"`
	// 已支付总金额
	TotalPayedAmount int64 `json:"total_payed_amount,omitempty" xml:"total_payed_amount,omitempty"`
	// 支付宝交易号
	AlipayOrderId string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`
	// 支付宝外部商户号
	AlipayOuterId string `json:"alipay_outer_id,omitempty" xml:"alipay_outer_id,omitempty"`
	// 扩展属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单ID
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 付款金额总额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 正向资金单ID，仅在逆向资金单上使用
	PayFundId int64 `json:"pay_fund_id,omitempty" xml:"pay_fund_id,omitempty"`
	// 支付订单号(即tp订单号或支付宝流水号)
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 付款方式
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 支付方账号
	PayerAccount string `json:"payer_account,omitempty" xml:"payer_account,omitempty"`
	// 支付方账号类型
	PayerAccountType int64 `json:"payer_account_type,omitempty" xml:"payer_account_type,omitempty"`
	// 付款方账号id
	PayerAlipayId string `json:"payer_alipay_id,omitempty" xml:"payer_alipay_id,omitempty"`
	// 支付方昵称
	PayerNick string `json:"payer_nick,omitempty" xml:"payer_nick,omitempty"`
	// 支付方淘宝ID
	PayerTid int64 `json:"payer_tid,omitempty" xml:"payer_tid,omitempty"`
	// 收款方账号
	ReceiverAccount string `json:"receiver_account,omitempty" xml:"receiver_account,omitempty"`
	// 收款方账号类型
	ReceiverAccountType int64 `json:"receiver_account_type,omitempty" xml:"receiver_account_type,omitempty"`
	// 收款方账号id
	ReceiverAlipayId string `json:"receiver_alipay_id,omitempty" xml:"receiver_alipay_id,omitempty"`
	// 收款方昵称
	ReceiverNick string `json:"receiver_nick,omitempty" xml:"receiver_nick,omitempty"`
	// 收款方淘宝ID
	ReceiverTid int64 `json:"receiver_tid,omitempty" xml:"receiver_tid,omitempty"`
	// 请求版本号
	ReqVersion string `json:"req_version,omitempty" xml:"req_version,omitempty"`
	// 二级商户smid
	Smid string `json:"smid,omitempty" xml:"smid,omitempty"`
	// 资金单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 付款来源
	TradeSource int64 `json:"trade_source,omitempty" xml:"trade_source,omitempty"`
	// 交易方式
	TradeType int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
}
