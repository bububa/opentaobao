package btrip

// BtripHotelValidateOrderRS 
type BtripHotelValidateOrderRS struct {
    // 创单Key值，下单时使用
    CreateKey   string `json:"create_key,omitempty" xml:"create_key,omitempty"`
    // 优惠信息
    PromotionInfo   *BtripHotelPromotionDTO `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
    // 销售计划id
    RatePlanId   int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
    // 详细的销售计划
    RatePlanInfo   *BtripHotelRatePlanInfoDTO `json:"rate_plan_info,omitempty" xml:"rate_plan_info,omitempty"`
}
