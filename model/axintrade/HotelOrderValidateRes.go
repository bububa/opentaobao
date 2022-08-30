package axintrade

// HotelOrderValidateRes 结构体
type HotelOrderValidateRes struct {
	// 创建订单需要的key
	CreateKey string `json:"create_key,omitempty" xml:"create_key,omitempty"`
	// 后返佣金费率
	CommissionFeeRate string `json:"commission_fee_rate,omitempty" xml:"commission_fee_rate,omitempty"`
	// 资源标签
	ResourceTag string `json:"resource_tag,omitempty" xml:"resource_tag,omitempty"`
	// 分销模式 FP-底价模式 SP-卖家模式
	DistributeMode string `json:"distribute_mode,omitempty" xml:"distribute_mode,omitempty"`
	// 售卖政策
	RatePlanInfo *RatePlanInfoApiDto `json:"rate_plan_info,omitempty" xml:"rate_plan_info,omitempty"`
	// 售卖政策id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 优惠总金额，单位为分
	PromotionTotalPrice int64 `json:"promotion_total_price,omitempty" xml:"promotion_total_price,omitempty"`
}
