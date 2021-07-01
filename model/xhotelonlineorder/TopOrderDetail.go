package xhotelonlineorder

// TopOrderDetail 结构体
type TopOrderDetail struct {
	// 最早到店时间
	ArriveEarly string `json:"arrive_early,omitempty" xml:"arrive_early,omitempty"`
	// 卖家没有确认直接关单剩余时间
	CloseOrderRestSecond int64 `json:"close_order_rest_second,omitempty" xml:"close_order_rest_second,omitempty"`
	// 飞猪酒店ID
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// rp名称
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 飞猪rpid
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
	// 完结时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 订单类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 订单关联包信息： topOrderPackageDO
	TopOrderPackage *TopOrderPackageDo `json:"top_order_package,omitempty" xml:"top_order_package,omitempty"`
	// 商家订单ID
	OutOid string `json:"out_oid,omitempty" xml:"out_oid,omitempty"`
	// 下单时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 商家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 酒店名称
	HotelName string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 住几晚
	Nights int64 `json:"nights,omitempty" xml:"nights,omitempty"`
	// 飞猪订单ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 订单垂直属性信息
	TopInfoMap *Topinfomap `json:"top_info_map,omitempty" xml:"top_info_map,omitempty"`
	// 外部订单的下单确认码
	OutConfirmCode string `json:"out_confirm_code,omitempty" xml:"out_confirm_code,omitempty"`
	// 交易状态
	TradeStatus string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
	// 支付宝交易流水号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 商家酒店Code
	HotelCode string `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
	// 商家rpcode
	RpCode string `json:"rp_code,omitempty" xml:"rp_code,omitempty"`
	// 是否包房
	ReservedRoomOrder bool `json:"reserved_room_order,omitempty" xml:"reserved_room_order,omitempty"`
	// 关联子订单对象
	TopRelationOrders []Toprelationorders `json:"top_relation_orders,omitempty" xml:"top_relation_orders>toprelationorders,omitempty"`
	// 是否拆单
	SplitOrder bool `json:"split_order,omitempty" xml:"split_order,omitempty"`
	// 酒店订单卖家信息
	HotelOrderSeller *HotelOrderSellerDo `json:"hotel_order_seller,omitempty" xml:"hotel_order_seller,omitempty"`
	// 联系人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 渠道企业名称
	OutSourceCorpName string `json:"out_source_corp_name,omitempty" xml:"out_source_corp_name,omitempty"`
	// 预定状态 面付信用住有效
	BookingStatus int64 `json:"booking_status,omitempty" xml:"booking_status,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商家房型Code
	RoomTypeCode string `json:"room_type_code,omitempty" xml:"room_type_code,omitempty"`
	// 消费码 全场景消费
	RelatedCode string `json:"related_code,omitempty" xml:"related_code,omitempty"`
	// 入住时间
	CheckinDate string `json:"checkin_date,omitempty" xml:"checkin_date,omitempty"`
	// List<TopDailyInfo>
	TopDailyInfos []TopDailyInfo `json:"top_daily_infos,omitempty" xml:"top_daily_infos>top_daily_info,omitempty"`
	// 房间数
	RoomNumber int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// 订单确认时长
	ConfirmDuration int64 `json:"confirm_duration,omitempty" xml:"confirm_duration,omitempty"`
	// 离店时间
	CheckoutDate string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	// 海外酒店信息
	TopOverseasPriceInfo *TopOverseasPriceInfoDo `json:"top_overseas_price_info,omitempty" xml:"top_overseas_price_info,omitempty"`
	// 是否需要确认号
	NeedConfirmNo bool `json:"need_confirm_no,omitempty" xml:"need_confirm_no,omitempty"`
	// 是否展示“卖家延迟按钮”
	ShowSellerDelayConfirmButton bool `json:"show_seller_delay_confirm_button,omitempty" xml:"show_seller_delay_confirm_button,omitempty"`
	// 最晚到店时间
	ArriveLate string `json:"arrive_late,omitempty" xml:"arrive_late,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 是否保留房
	BlockRoomOrder bool `json:"block_room_order,omitempty" xml:"block_room_order,omitempty"`
	// 用户支付费用
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 订单取消原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 飞猪房型ID
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 商家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 子订单类型
	SubType int64 `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 总房费
	TotalRoomPrice int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
	// outSource
	OutSource string `json:"out_source,omitempty" xml:"out_source,omitempty"`
	// 房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	// List<TopOrderGuest>
	TopOrderGuests []TopOrderGuest `json:"top_order_guests,omitempty" xml:"top_order_guests>top_order_guest,omitempty"`
	// 统一订单状态，与union_status_value相对应。101:已下单;102:已付款;103:确认有房;104:交易成功;105:申请退款;106:卖家拒绝退款;108:订单关闭;109:退款成功;522:买家待支付;501:已下单;502:确认有房;503:确认前用户取消;504:确认无房;505:买家入住;506:系统自动核实入住成功;507:确认后用户取消;508:卖家核实未入住NoShow;601:已下单;602:确认有房;603:确认前用户取消;604:确认无房;605:买家入住;606:系统自动核实入住成功;607:确认后用户取消;608:卖家核实未入住NoShow;609:买家离店;610:结账中;611:结账成功;612:结账失败追账中;613:追账成功;614:NOSHOW结账中;615:NOSHOW结账成功;616:NOSHOW结账失败;617:NOSHOW追账成功;618:结账取消;619:系统自动核实未入住成功;620:结账失败追账失败;621:结账失败追账成功;701:已下单;702:确认有房;703:用户取消;704:确认无房;705:买家入住;709:买家离店;710:取消中
	UnionStatusText string `json:"union_status_text,omitempty" xml:"union_status_text,omitempty"`
	// 统一订单状态
	UnionStatusValue int64 `json:"union_status_value,omitempty" xml:"union_status_value,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 其他费用
	OtherFee int64 `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
	// memo
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 物流状态描述
	LogisticsStatusText string `json:"logistics_status_text,omitempty" xml:"logistics_status_text,omitempty"`
	// 物流状态值
	LogisticsStatusValue int64 `json:"logistics_status_value,omitempty" xml:"logistics_status_value,omitempty"`
	// 支付状态
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 担保类型
	GuaranteeType int64 `json:"guarantee_type,omitempty" xml:"guarantee_type,omitempty"`
	// 担保id
	FundId int64 `json:"fund_id,omitempty" xml:"fund_id,omitempty"`
	// 取消规则描述
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// topPromotion
	TopPromotion *TopPromotion `json:"top_promotion,omitempty" xml:"top_promotion,omitempty"`
	// 是否即时确认
	TagJsqr bool `json:"tag_jsqr,omitempty" xml:"tag_jsqr,omitempty"`
	// 房价id
	RateId string `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 税和费，单位分
	TaxAndFee int64 `json:"tax_and_fee,omitempty" xml:"tax_and_fee,omitempty"`
	// topDomesticPriceInfoDO
	TopDomesticPriceInfo *TopDomesticPriceInfoDo `json:"top_domestic_price_info,omitempty" xml:"top_domestic_price_info,omitempty"`
	// 用户实际支付金额
	UserRealPayment int64 `json:"user_real_payment,omitempty" xml:"user_real_payment,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// vendor
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 信用卡vcc信息
	CreditCard *CreditCard `json:"credit_card,omitempty" xml:"credit_card,omitempty"`
	// 退款状态，1：买家已经申请退款，等待卖家同意;4:退款关闭;5:退款成功;6:卖家拒绝退款;9:没有申请退款
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款金额  单位为分
	RefundAmout int64 `json:"refund_amout,omitempty" xml:"refund_amout,omitempty"`
	// 退款信息
	PostTradeRefund *TopPostTradeRefundDo `json:"post_trade_refund,omitempty" xml:"post_trade_refund,omitempty"`
	// 销售渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 买家id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
	// 城市code
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 支付类型
	PaymentType string `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 发票状态
	InvoiceStatus string `json:"invoice_status,omitempty" xml:"invoice_status,omitempty"`
	// 订单来源类型
	OrderSourceType string `json:"order_source_type,omitempty" xml:"order_source_type,omitempty"`
	// topInvoiceDO
	TopInvoice *TopInvoiceDo `json:"top_invoice,omitempty" xml:"top_invoice,omitempty"`
	// 退款原因类型如101，503之类
	RefundReasonType int64 `json:"refund_reason_type,omitempty" xml:"refund_reason_type,omitempty"`
}
