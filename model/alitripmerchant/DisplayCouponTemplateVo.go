package alitripmerchant

// DisplayCouponTemplateVo 结构体
type DisplayCouponTemplateVo struct {
	// 关联活动id
	RelateActivityIdList []int64 `json:"relate_activity_id_list,omitempty" xml:"relate_activity_id_list>int64,omitempty"`
	// 使用门槛
	ConditionAmount string `json:"condition_amount,omitempty" xml:"condition_amount,omitempty"`
	// 面额
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 优惠券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 使用条件描述
	ConditionDesc string `json:"condition_desc,omitempty" xml:"condition_desc,omitempty"`
	// 详细描述
	DetailDesc string `json:"detail_desc,omitempty" xml:"detail_desc,omitempty"`
	// 优惠券预订开始时间
	BookStartTime string `json:"book_start_time,omitempty" xml:"book_start_time,omitempty"`
	// 优惠券预订结束时间
	BookEndTime string `json:"book_end_time,omitempty" xml:"book_end_time,omitempty"`
	// 允许入住开始时间
	CheckInTime string `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// 允许入住结束时间
	CheckOutTime string `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// 关联状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 优惠券过期时间
	ExpiredTime string `json:"expired_time,omitempty" xml:"expired_time,omitempty"`
	// 优惠券类型
	CouponType string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 优惠券模板id
	CouponTemplateId int64 `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
	// 领取优惠券后的过期时间
	ExpiredTimeDay int64 `json:"expired_time_day,omitempty" xml:"expired_time_day,omitempty"`
	// 领取优惠券后的过期时间
	ExpiredTimeHour int64 `json:"expired_time_hour,omitempty" xml:"expired_time_hour,omitempty"`
	// 领取优惠券后的过期时间
	ExpiredTimeMin int64 `json:"expired_time_min,omitempty" xml:"expired_time_min,omitempty"`
	// 优惠券总数量
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 已领取数量
	ReceivedAmount int64 `json:"received_amount,omitempty" xml:"received_amount,omitempty"`
	// 优惠券剩余数量
	LeftAmount int64 `json:"left_amount,omitempty" xml:"left_amount,omitempty"`
	// 关联商品
	ReplaceRpCode *ReplaceRpCode `json:"replace_rp_code,omitempty" xml:"replace_rp_code,omitempty"`
}
