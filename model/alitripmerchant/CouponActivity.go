package alitripmerchant

import (
	"sync"
)

// CouponActivity 结构体
type CouponActivity struct {
	// 优惠券详情
	CouponList []CouponInfo `json:"coupon_list,omitempty" xml:"coupon_list>coupon_info,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolCouponActivity = sync.Pool{
	New: func() any {
		return new(CouponActivity)
	},
}

// GetCouponActivity() 从对象池中获取CouponActivity
func GetCouponActivity() *CouponActivity {
	return poolCouponActivity.Get().(*CouponActivity)
}

// ReleaseCouponActivity 释放CouponActivity
func ReleaseCouponActivity(v *CouponActivity) {
	v.CouponList = v.CouponList[:0]
	v.ActivityName = ""
	v.ActivityId = 0
	poolCouponActivity.Put(v)
}
