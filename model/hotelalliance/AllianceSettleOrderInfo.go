package hotelalliance

// AllianceSettleOrderInfo 结构体
type AllianceSettleOrderInfo struct {
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 订单渠道来源
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 分账日期
	SettleDate string `json:"settle_date,omitempty" xml:"settle_date,omitempty"`
	// 分账费率
	SettleRate string `json:"settle_rate,omitempty" xml:"settle_rate,omitempty"`
	// 分账状态
	SettleStatus int64 `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 杂费（单位：分）
	OtherFee int64 `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
	// 分账金额（单位：分）
	SettleAmount int64 `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	// 卖家佣金金额（单位：分）
	SellerCommission int64 `json:"seller_commission,omitempty" xml:"seller_commission,omitempty"`
	// 飞猪订单号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 分账拦截金额（单位：分）
	InterceptAmount int64 `json:"intercept_amount,omitempty" xml:"intercept_amount,omitempty"`
	// 卖家Id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 是否底价订单
	BaseMode int64 `json:"base_mode,omitempty" xml:"base_mode,omitempty"`
	// 订单有效间夜
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 房费（单位：分）
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 底价金额（单位：分）
	BasePrice int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
}
