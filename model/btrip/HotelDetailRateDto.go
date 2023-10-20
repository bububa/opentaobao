package btrip

import (
	"sync"
)

// HotelDetailRateDto 结构体
type HotelDetailRateDto struct {
	// 每日优惠后价格DTO
	PricePlan []DiscountDailyPrice `json:"price_plan,omitempty" xml:"price_plan>discount_daily_price,omitempty"`
	// 早餐描述
	Breakfast string `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	// 免费取消政策描述
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// 协议价，1表示协议支付
	CompanyAassist string `json:"company_aassist,omitempty" xml:"company_aassist,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 每日截止时间，可忽略
	EndTimeDaily string `json:"end_time_daily,omitempty" xml:"end_time_daily,omitempty"`
	// 库存价格
	InventoryPrice string `json:"inventory_price,omitempty" xml:"inventory_price,omitempty"`
	// 营销信息
	PromotionInfo string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// 报价名称
	RatePlanName string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty"`
	// 每日开始时间。可忽略
	StartTimeDaily string `json:"start_time_daily,omitempty" xml:"start_time_daily,omitempty"`
	// 供应商标识码
	SupplierCode string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// 供应商名称
	SupplierName string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 最小提前小时数，可忽略
	MinAdvHours int64 `json:"min_adv_hours,omitempty" xml:"min_adv_hours,omitempty"`
	// 最小间隔天数，可忽略
	MinDays int64 `json:"min_days,omitempty" xml:"min_days,omitempty"`
	// 间隔天数。可忽略
	Nod int64 `json:"nod,omitempty" xml:"nod,omitempty"`
	// 人数
	Nop int64 `json:"nop,omitempty" xml:"nop,omitempty"`
	// 支付方式，1-全额支付
	PaymentType int64 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// 报价ID
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// rpId
	RpId int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// 取消政策
	BtripHotelCancelPolicyDTO *BtripHotelCancelPolicyDto `json:"btrip_hotel_cancel_policy_d_t_o,omitempty" xml:"btrip_hotel_cancel_policy_d_t_o,omitempty"`
	// 取整后每日均价
	DailyPriceFormatYuan int64 `json:"daily_price_format_yuan,omitempty" xml:"daily_price_format_yuan,omitempty"`
	// 界面展示价格
	DailyPriceView int64 `json:"daily_price_view,omitempty" xml:"daily_price_view,omitempty"`
	// 是否即时确认，可忽略
	InstantConfirm bool `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
	// 是否支持专票
	SupportSpecialInvoice bool `json:"support_special_invoice,omitempty" xml:"support_special_invoice,omitempty"`
}

var poolHotelDetailRateDto = sync.Pool{
	New: func() any {
		return new(HotelDetailRateDto)
	},
}

// GetHotelDetailRateDto() 从对象池中获取HotelDetailRateDto
func GetHotelDetailRateDto() *HotelDetailRateDto {
	return poolHotelDetailRateDto.Get().(*HotelDetailRateDto)
}

// ReleaseHotelDetailRateDto 释放HotelDetailRateDto
func ReleaseHotelDetailRateDto(v *HotelDetailRateDto) {
	v.PricePlan = v.PricePlan[:0]
	v.Breakfast = ""
	v.CancelPolicyDesc = ""
	v.CompanyAassist = ""
	v.CurrencyCode = ""
	v.EndTimeDaily = ""
	v.InventoryPrice = ""
	v.PromotionInfo = ""
	v.RatePlanName = ""
	v.StartTimeDaily = ""
	v.SupplierCode = ""
	v.SupplierName = ""
	v.ItemId = 0
	v.MinAdvHours = 0
	v.MinDays = 0
	v.Nod = 0
	v.Nop = 0
	v.PaymentType = 0
	v.RateId = 0
	v.RpId = 0
	v.BtripHotelCancelPolicyDTO = nil
	v.DailyPriceFormatYuan = 0
	v.DailyPriceView = 0
	v.InstantConfirm = false
	v.SupportSpecialInvoice = false
	poolHotelDetailRateDto.Put(v)
}
