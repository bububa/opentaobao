package alitripmerchant

// CouponActivity 结构体
type CouponActivity struct {
	// 优惠券详情
	CouponList []CouponInfo `json:"coupon_list,omitempty" xml:"coupon_list>coupon_info,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
