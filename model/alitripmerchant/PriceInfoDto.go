package alitripmerchant

import (
	"sync"
)

// PriceInfoDto 结构体
type PriceInfoDto struct {
	// 早餐信息 -1 含早、0无早、1单早、2、双早
	Breakfasts []string `json:"breakfasts,omitempty" xml:"breakfasts>string,omitempty"`
	// 详情页会员权益标签
	HotelDinamicLabels []HotelDinamicLabelDto `json:"hotel_dinamic_labels,omitempty" xml:"hotel_dinamic_labels>hotel_dinamic_label_dto,omitempty"`
	// 绑定优惠券模板id
	CouponTemplateIds []int64 `json:"coupon_template_ids,omitempty" xml:"coupon_template_ids>int64,omitempty"`
	// 税费 单位元
	TaxAndFeeToString string `json:"tax_and_fee_to_string,omitempty" xml:"tax_and_fee_to_string,omitempty"`
	// 含税价 均价 单位元
	DinamicOriginalPriceToString string `json:"dinamic_original_price_to_string,omitempty" xml:"dinamic_original_price_to_string,omitempty"`
	// 退订规则
	RefundRules string `json:"refund_rules,omitempty" xml:"refund_rules,omitempty"`
	// 可够取消险
	CancelInsuranceIcon string `json:"cancel_insurance_icon,omitempty" xml:"cancel_insurance_icon,omitempty"`
	// 净价均价 不含税 单位元
	RatePriceToString string `json:"rate_price_to_string,omitempty" xml:"rate_price_to_string,omitempty"`
	// 会员等级
	MemberLevel string `json:"member_level,omitempty" xml:"member_level,omitempty"`
	// 币种CNY 人民币
	PriceType string `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 早餐
	BreakfastDesc string `json:"breakfast_desc,omitempty" xml:"breakfast_desc,omitempty"`
	// 取消政策模型同小搜一致即可，入住日往前XX小时
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 不含税总价 保留到元
	RoomTotalPriceToString string `json:"room_total_price_to_string,omitempty" xml:"room_total_price_to_string,omitempty"`
	// 取消险文案
	CancelInsuranceDesc string `json:"cancel_insurance_desc,omitempty" xml:"cancel_insurance_desc,omitempty"`
	// 商品定价规则标题
	RpName string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
	// 库存数量
	StockNumberDes string `json:"stock_number_des,omitempty" xml:"stock_number_des,omitempty"`
	// 首住标签
	FirstLive string `json:"first_live,omitempty" xml:"first_live,omitempty"`
	// 副会员商品等级
	MemberLevelV2 string `json:"member_level_v2,omitempty" xml:"member_level_v2,omitempty"`
	// rpcode
	RpCode string `json:"rp_code,omitempty" xml:"rp_code,omitempty"`
	// 标签
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
	// 优惠券类型
	CouponType string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 权益券类型
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// APLUSP
	VoucherCardCategory string `json:"voucher_card_category,omitempty" xml:"voucher_card_category,omitempty"`
	// 划线价
	UnderscorePrice string `json:"underscore_price,omitempty" xml:"underscore_price,omitempty"`
	// 是否担保商品
	IsGuarantee int64 `json:"is_guarantee,omitempty" xml:"is_guarantee,omitempty"`
	// 库存剩余数量
	StockNumber int64 `json:"stock_number,omitempty" xml:"stock_number,omitempty"`
	// 商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 取消政策枚举 01-不可取消 02-免费取消 03-阶梯取消 04-限时免费取消
	CancelType int64 `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
	// 含税价 均价 单位分
	DinamicOriginalPrice int64 `json:"dinamic_original_price,omitempty" xml:"dinamic_original_price,omitempty"`
	// 是否已售空
	IsSellOut int64 `json:"is_sell_out,omitempty" xml:"is_sell_out,omitempty"`
	// 不含税总价 单位分
	RoomTotalPrice int64 `json:"room_total_price,omitempty" xml:"room_total_price,omitempty"`
	// Rate Plan ID商品定价规则
	RpId int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// 商品最小可售卖粒度
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 1:全额支付;5:前台面付;
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 净价均价 不含税
	RatePrice int64 `json:"rate_price,omitempty" xml:"rate_price,omitempty"`
	// 税费 单位分
	TaxAndFee int64 `json:"tax_and_fee,omitempty" xml:"tax_and_fee,omitempty"`
	// 卖家房型id
	OuterRoomId int64 `json:"outer_room_id,omitempty" xml:"outer_room_id,omitempty"`
	// 表示当前价格允许几人入住
	AllowPersonNumber int64 `json:"allow_person_number,omitempty" xml:"allow_person_number,omitempty"`
	// 是否售空
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 房型id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 热门商品权重
	HotPrice int64 `json:"hot_price,omitempty" xml:"hot_price,omitempty"`
	// 折扣商品
	PriceInfo *PriceInfoDto `json:"price_info,omitempty" xml:"price_info,omitempty"`
	// 折扣总价
	DiscountAmount int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 代金券信息
	VoucherInfo *VoucherVo `json:"voucher_info,omitempty" xml:"voucher_info,omitempty"`
	// 折扣百分比
	DiscountOff int64 `json:"discount_off,omitempty" xml:"discount_off,omitempty"`
	// 0 立即预订/1 前去购买,
	BookingOrBuy int64 `json:"booking_or_buy,omitempty" xml:"booking_or_buy,omitempty"`
	// 是否是会员房
	MemberRoom bool `json:"member_room,omitempty" xml:"member_room,omitempty"`
	// 是否是副会员房
	MemberRoomV2 bool `json:"member_room_v2,omitempty" xml:"member_room_v2,omitempty"`
	// 是否为权益商品房型
	IsDerbyVoucherRoom bool `json:"is_derby_voucher_room,omitempty" xml:"is_derby_voucher_room,omitempty"`
}

var poolPriceInfoDto = sync.Pool{
	New: func() any {
		return new(PriceInfoDto)
	},
}

// GetPriceInfoDto() 从对象池中获取PriceInfoDto
func GetPriceInfoDto() *PriceInfoDto {
	return poolPriceInfoDto.Get().(*PriceInfoDto)
}

// ReleasePriceInfoDto 释放PriceInfoDto
func ReleasePriceInfoDto(v *PriceInfoDto) {
	v.Breakfasts = v.Breakfasts[:0]
	v.HotelDinamicLabels = v.HotelDinamicLabels[:0]
	v.CouponTemplateIds = v.CouponTemplateIds[:0]
	v.TaxAndFeeToString = ""
	v.DinamicOriginalPriceToString = ""
	v.RefundRules = ""
	v.CancelInsuranceIcon = ""
	v.RatePriceToString = ""
	v.MemberLevel = ""
	v.PriceType = ""
	v.BreakfastDesc = ""
	v.CancelTime = ""
	v.RoomTotalPriceToString = ""
	v.CancelInsuranceDesc = ""
	v.RpName = ""
	v.StockNumberDes = ""
	v.FirstLive = ""
	v.MemberLevelV2 = ""
	v.RpCode = ""
	v.Tag = ""
	v.CouponType = ""
	v.Category = ""
	v.VoucherCardCategory = ""
	v.UnderscorePrice = ""
	v.IsGuarantee = 0
	v.StockNumber = 0
	v.Gid = 0
	v.CancelType = 0
	v.DinamicOriginalPrice = 0
	v.IsSellOut = 0
	v.RoomTotalPrice = 0
	v.RpId = 0
	v.RateId = 0
	v.PaymentType = 0
	v.RatePrice = 0
	v.TaxAndFee = 0
	v.OuterRoomId = 0
	v.AllowPersonNumber = 0
	v.Status = 0
	v.RoomId = 0
	v.HotPrice = 0
	v.PriceInfo = nil
	v.DiscountAmount = 0
	v.VoucherInfo = nil
	v.DiscountOff = 0
	v.BookingOrBuy = 0
	v.MemberRoom = false
	v.MemberRoomV2 = false
	v.IsDerbyVoucherRoom = false
	poolPriceInfoDto.Put(v)
}
