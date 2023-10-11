package wdk

// OrderActivity 结构体
type OrderActivity struct {
	// 渠道活动ID
	ChannelActivityId string `json:"channel_activity_id,omitempty" xml:"channel_activity_id,omitempty"`
	// 关联业务活动ID
	BizActivityId string `json:"biz_activity_id,omitempty" xml:"biz_activity_id,omitempty"`
	// 关联商家ERP活动ID
	MerchantActivityId string `json:"merchant_activity_id,omitempty" xml:"merchant_activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动类型
	ActivityType string `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 活动优惠商家分摊
	DiscountMerchantFee int64 `json:"discount_merchant_fee,omitempty" xml:"discount_merchant_fee,omitempty"`
	// 活动优惠平台分摊
	DiscountPlatformFee int64 `json:"discount_platform_fee,omitempty" xml:"discount_platform_fee,omitempty"`
}
