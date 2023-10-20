package promotion

import (
	"sync"
)

// CouponActivity 结构体
type CouponActivity struct {
	// 匿名码code
	MaCode string `json:"ma_code,omitempty" xml:"ma_code,omitempty"`
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
	v.MaCode = ""
	poolCouponActivity.Put(v)
}
