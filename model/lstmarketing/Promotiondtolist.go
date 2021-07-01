package lstmarketing

// Promotiondtolist 结构体
type Promotiondtolist struct {
	// 优惠金额，分为单位
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 营销类型
	PromotionTypeName string `json:"promotion_type_name,omitempty" xml:"promotion_type_name,omitempty"`
}
