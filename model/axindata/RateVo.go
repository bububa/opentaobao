package axindata

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
	// 是否即时确认
	InstantConfirm bool `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
}
