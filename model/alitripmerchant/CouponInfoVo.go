package alitripmerchant

// CouponInfoVo 结构体
type CouponInfoVo struct {
	// 优惠券状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 过期时间
	ExpiredTimeMin string `json:"expired_time_min,omitempty" xml:"expired_time_min,omitempty"`
	// 详细描述
	DetailDesc string `json:"detail_desc,omitempty" xml:"detail_desc,omitempty"`
	// 使用条件描述
	ConditionDesc string `json:"condition_desc,omitempty" xml:"condition_desc,omitempty"`
	// 优惠券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 使用门槛
	ConditionAmount string `json:"condition_amount,omitempty" xml:"condition_amount,omitempty"`
	// 面额(减免额度)
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 折扣券
	CouponType string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 优惠券实例id
	InstanceId int64 `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
}
