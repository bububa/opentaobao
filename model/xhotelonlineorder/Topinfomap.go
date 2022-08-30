package xhotelonlineorder

// Topinfomap 结构体
type Topinfomap struct {
	// hbsEgOrder
	HbsEgOrder string `json:"hbs_eg_order,omitempty" xml:"hbs_eg_order,omitempty"`
	// hbsBuyerAlipayNo
	HbsBuyerAlipayNo string `json:"hbs_buyer_alipay_no,omitempty" xml:"hbs_buyer_alipay_no,omitempty"`
	// preChangePromotionFlag
	PreChangePromotionFlag string `json:"pre_change_promotion_flag,omitempty" xml:"pre_change_promotion_flag,omitempty"`
	// 凌晨入住订单
	MorningBuy string `json:"morning_buy,omitempty" xml:"morning_buy,omitempty"`
	// cancelType
	CancelType string `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 卖家延迟确认时间，单位分钟
	SellerDelayConfirmTime string `json:"seller_delay_confirm_time,omitempty" xml:"seller_delay_confirm_time,omitempty"`
	// currencyType
	CurrencyType string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
	// hbsBedType
	HbsBedType string `json:"hbs_bed_type,omitempty" xml:"hbs_bed_type,omitempty"`
	// umid
	Umid string `json:"umid,omitempty" xml:"umid,omitempty"`
	// hbsOldPromotionamt
	HbsOldPromotionamt string `json:"hbs_old_promotionamt,omitempty" xml:"hbs_old_promotionamt,omitempty"`
	// hbsMaxOtherFee
	HbsMaxOtherFee string `json:"hbs_max_other_fee,omitempty" xml:"hbs_max_other_fee,omitempty"`
	// cancelPolicyJson
	CancelPolicyJson string `json:"cancel_policy_json,omitempty" xml:"cancel_policy_json,omitempty"`
	// hbsAlipayMobile
	HbsAlipayMobile string `json:"hbs_alipay_mobile,omitempty" xml:"hbs_alipay_mobile,omitempty"`
	// pms订单号
	HbsPmsOrderId string `json:"hbs_pms_order_id,omitempty" xml:"hbs_pms_order_id,omitempty"`
	// hbsHotelOutId
	HbsHotelOutId string `json:"hbs_hotel_out_id,omitempty" xml:"hbs_hotel_out_id,omitempty"`
	// hbsCreditSettle
	HbsCreditSettle string `json:"hbs_credit_settle,omitempty" xml:"hbs_credit_settle,omitempty"`
	// ebookingDirectTag
	EbookingDirectTag string `json:"ebooking_direct_tag,omitempty" xml:"ebooking_direct_tag,omitempty"`
	// hbsOldPayment
	HbsOldPayment string `json:"hbs_old_payment,omitempty" xml:"hbs_old_payment,omitempty"`
	// 身份证线下信用住
	HbsCardAlipayOrder string `json:"hbs_card_alipay_order,omitempty" xml:"hbs_card_alipay_order,omitempty"`
	// promotionSnapKey
	PromotionSnapKey string `json:"promotion_snap_key,omitempty" xml:"promotion_snap_key,omitempty"`
	// hbsAllTags
	HbsAllTags string `json:"hbs_all_tags,omitempty" xml:"hbs_all_tags,omitempty"`
	// 是否直连订单
	HbsDirect string `json:"hbs_direct,omitempty" xml:"hbs_direct,omitempty"`
	// 底价加价
	UpsetPrice string `json:"upset_price,omitempty" xml:"upset_price,omitempty"`
	// 是否包房
	ReservedRoom string `json:"reserved_room,omitempty" xml:"reserved_room,omitempty"`
	// 是否保留房
	BlockRoom string `json:"block_room,omitempty" xml:"block_room,omitempty"`
	// 渠道企业名称
	HbsOutSourceCorpName string `json:"hbs_out_source_corp_name,omitempty" xml:"hbs_out_source_corp_name,omitempty"`
	// 扫码信用住
	HbsScanCodeOrder string `json:"hbs_scan_code_order,omitempty" xml:"hbs_scan_code_order,omitempty"`
	// 商旅企业支付订单标识，1为是
	BtripCorporatePay string `json:"btrip_corporate_pay,omitempty" xml:"btrip_corporate_pay,omitempty"`
	// 是否是在线预约订单
	OnlineBookingRoom string `json:"online_booking_room,omitempty" xml:"online_booking_room,omitempty"`
	// b2g标识
	B2gFlag string `json:"b2g_flag,omitempty" xml:"b2g_flag,omitempty"`
	// hbsSelfHelpCheckIn
	HbsSelfHelpCheckIn string `json:"hbs_self_help_check_in,omitempty" xml:"hbs_self_help_check_in,omitempty"`
	// 小时房标识 1:小时房
	HbsXiaoShiFang string `json:"hbs_xiao_shi_fang,omitempty" xml:"hbs_xiao_shi_fang,omitempty"`
	// 小时房时长
	HbsHourage string `json:"hbs_hourage,omitempty" xml:"hbs_hourage,omitempty"`
	// 预约发票标识
	AdvanceInvocieTag string `json:"advance_invocie_tag,omitempty" xml:"advance_invocie_tag,omitempty"`
	// 小时房可入住时间
	HourRoomReservedTime string `json:"hour_room_reserved_time,omitempty" xml:"hour_room_reserved_time,omitempty"`
	// 小时房最晚离店时间
	HourRoomLeaveTime string `json:"hour_room_leave_time,omitempty" xml:"hour_room_leave_time,omitempty"`
	// 小时房到店时间
	HourRoomArriveTime string `json:"hour_room_arrive_time,omitempty" xml:"hour_room_arrive_time,omitempty"`
	// 信用住升级订单
	CreditPayOrder string `json:"credit_pay_order,omitempty" xml:"credit_pay_order,omitempty"`
	// 是否是菲住联盟会员 true or false
	IsFzMemberOrder string `json:"is_fz_member_order,omitempty" xml:"is_fz_member_order,omitempty"`
	// 会员卡卡号
	BookerMemberCardNo string `json:"booker_member_card_no,omitempty" xml:"booker_member_card_no,omitempty"`
	// 渠道企业id
	HbsOutSourceCorpId string `json:"hbs_out_source_corp_id,omitempty" xml:"hbs_out_source_corp_id,omitempty"`
	// 飞猪会员等级
	FliggyMemberLevel string `json:"fliggy_member_level,omitempty" xml:"fliggy_member_level,omitempty"`
	// 是否是菲住会员
	IsFzMember string `json:"is_fz_member,omitempty" xml:"is_fz_member,omitempty"`
	// 合并支付会员卡订单号
	MemberCardOrderId string `json:"member_card_order_id,omitempty" xml:"member_card_order_id,omitempty"`
	// 订单间夜优惠明细,其中其中day:日期，sellerFullReduce:当日满减卖家优惠，sellerOtherReduce:当日非满减卖家优惠，sellerFullReduceAll:总满减卖家优惠，sellerOtherReduceAll:总非满减卖家优惠
	NightPromotionDetail string `json:"night_promotion_detail,omitempty" xml:"night_promotion_detail,omitempty"`
	// 信用支付因取消的扣款费用的人民币，单位：分
	CreditPayCancelFee int64 `json:"credit_pay_cancel_fee,omitempty" xml:"credit_pay_cancel_fee,omitempty"`
	// 售卖会员价
	SellingMemberPrice int64 `json:"selling_member_price,omitempty" xml:"selling_member_price,omitempty"`
	// 售卖价
	SellingPrice int64 `json:"selling_price,omitempty" xml:"selling_price,omitempty"`
}
