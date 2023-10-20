package axindata

import (
	"sync"
)

// RateVo 结构体
type RateVo struct {
	// 售卖政策名称
	RatePlanName string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty"`
	// 每日截止时间
	EndTimeDaily string `json:"end_time_daily,omitempty" xml:"end_time_daily,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 退改政策描述
	CancelPolicyDesc string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// 是否含早餐
	Breakfast string `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	// 佣金费率
	CommissionFeeRate string `json:"commission_fee_rate,omitempty" xml:"commission_fee_rate,omitempty"`
	// 资源标签
	ResourceTag string `json:"resource_tag,omitempty" xml:"resource_tag,omitempty"`
	// 分销模式 FP-底价模式 SP-卖家模式
	DistributeMode string `json:"distribute_mode,omitempty" xml:"distribute_mode,omitempty"`
	// 售卖政策id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 售卖单位id
	RateId int64 `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	// 最小提前小时数
	MinAdvHours int64 `json:"min_adv_hours,omitempty" xml:"min_adv_hours,omitempty"`
	// 库存
	Quota int64 `json:"quota,omitempty" xml:"quota,omitempty"`
	// 价格(单位分)
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 优惠金额（单位分）
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 取消政策
	CancelPolicyVO *CancelPolicyVo `json:"cancel_policy_v_o,omitempty" xml:"cancel_policy_v_o,omitempty"`
	// 早餐数量
	BreakfastCount int64 `json:"breakfast_count,omitempty" xml:"breakfast_count,omitempty"`
	// 凌晨房信息
	DawnBookingVo *DawnBookingVo `json:"dawn_booking_vo,omitempty" xml:"dawn_booking_vo,omitempty"`
	// 最小连住天数
	MinDays int64 `json:"min_days,omitempty" xml:"min_days,omitempty"`
	// 最大连住天数
	MaxDays int64 `json:"max_days,omitempty" xml:"max_days,omitempty"`
	// 最大提前小时数
	MaxAdvHours int64 `json:"max_adv_hours,omitempty" xml:"max_adv_hours,omitempty"`
	// 小时房信息
	HourRoomInfo *HourRoomInfo `json:"hour_room_info,omitempty" xml:"hour_room_info,omitempty"`
	// 小时房信息
	HourRoomInfoDto *HourRoomInfoDto `json:"hour_room_info_dto,omitempty" xml:"hour_room_info_dto,omitempty"`
	// 1-订单金额酒店开票，分销商开返佣发票 2-订单金额阿信开票，分销商不开票 3-订单金额无票，分销商开返佣发票
	InvoicingMode int64 `json:"invoicing_mode,omitempty" xml:"invoicing_mode,omitempty"`
	// 是否即时确认
	InstantConfirm bool `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
	// 是否小时房,不为空且为true时标识小时房，否则全日房
	HourRoom bool `json:"hour_room,omitempty" xml:"hour_room,omitempty"`
}

var poolRateVo = sync.Pool{
	New: func() any {
		return new(RateVo)
	},
}

// GetRateVo() 从对象池中获取RateVo
func GetRateVo() *RateVo {
	return poolRateVo.Get().(*RateVo)
}

// ReleaseRateVo 释放RateVo
func ReleaseRateVo(v *RateVo) {
	v.RatePlanName = ""
	v.EndTimeDaily = ""
	v.CurrencyCode = ""
	v.CancelPolicyDesc = ""
	v.Breakfast = ""
	v.CommissionFeeRate = ""
	v.ResourceTag = ""
	v.DistributeMode = ""
	v.RatePlanId = 0
	v.ItemId = 0
	v.RateId = 0
	v.MinAdvHours = 0
	v.Quota = 0
	v.Price = 0
	v.PromotionPrice = 0
	v.CancelPolicyVO = nil
	v.BreakfastCount = 0
	v.DawnBookingVo = nil
	v.MinDays = 0
	v.MaxDays = 0
	v.MaxAdvHours = 0
	v.HourRoomInfo = nil
	v.HourRoomInfoDto = nil
	v.InvoicingMode = 0
	v.InstantConfirm = false
	v.HourRoom = false
	poolRateVo.Put(v)
}
