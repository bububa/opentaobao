package c2m

// CompanyOrderInfoVo 结构体
type CompanyOrderInfoVo struct {
	// 商品标题
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 加密后的分销商上级ID
	EncryptInvitorId string `json:"encrypt_invitor_id,omitempty" xml:"encrypt_invitor_id,omitempty"`
	// 加密后的买家ID
	EncryptBuyerId string `json:"encrypt_buyer_id,omitempty" xml:"encrypt_buyer_id,omitempty"`
	// 加密后的分销商ID
	EncryptDistributorId string `json:"encrypt_distributor_id,omitempty" xml:"encrypt_distributor_id,omitempty"`
	// 订单修改日期
	ModifyDate string `json:"modify_date,omitempty" xml:"modify_date,omitempty"`
	// 订单确认收货日期
	SuccDate string `json:"succ_date,omitempty" xml:"succ_date,omitempty"`
	// 加密后的订单商品标
	EncryptItemTag string `json:"encrypt_item_tag,omitempty" xml:"encrypt_item_tag,omitempty"`
	// 收货人区域
	MordArea string `json:"mord_area,omitempty" xml:"mord_area,omitempty"`
	// 收货人城市
	MordCity string `json:"mord_city,omitempty" xml:"mord_city,omitempty"`
	// 收货人省份
	MordProv string `json:"mord_prov,omitempty" xml:"mord_prov,omitempty"`
	// 导师名称
	TutorNick string `json:"tutor_nick,omitempty" xml:"tutor_nick,omitempty"`
	// 班级名称
	ClassName string `json:"class_name,omitempty" xml:"class_name,omitempty"`
	// 上级导师名称
	UpTutorNick string `json:"up_tutor_nick,omitempty" xml:"up_tutor_nick,omitempty"`
	// 佣金金额（单位：分）
	CommissionFee int64 `json:"commission_fee,omitempty" xml:"commission_fee,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 付款金额（单位：分）
	TotalFee int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 商品ID
	AuctionId int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 订单类型，0-自购，1-代下单，2-分享购买
	Delegate int64 `json:"delegate,omitempty" xml:"delegate,omitempty"`
	// 子订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 商品类型，0-普通商品，1-高佣商品
	AuctionType int64 `json:"auction_type,omitempty" xml:"auction_type,omitempty"`
	// 退款金额（单位：分）
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 订单付款时间
	PayTime int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 订单创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 班级ID
	ClassId int64 `json:"class_id,omitempty" xml:"class_id,omitempty"`
	// 订单规则ID
	RuleId int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
}
