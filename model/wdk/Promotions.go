package wdk

// Promotions 结构体
type Promotions struct {
	// 活动优惠金额
	DiscountFee string `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动id
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
