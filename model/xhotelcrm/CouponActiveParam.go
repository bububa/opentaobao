package xhotelcrm

import (
	"sync"
)

// CouponActiveParam 结构体
type CouponActiveParam struct {
	// 券有效期截止时间
	CouponEndDate string `json:"coupon_end_date,omitempty" xml:"coupon_end_date,omitempty"`
	// 券实例ID
	CouponInstanceId string `json:"coupon_instance_id,omitempty" xml:"coupon_instance_id,omitempty"`
	// 券有效期考试时间
	CouponStartDate string `json:"coupon_start_date,omitempty" xml:"coupon_start_date,omitempty"`
}

var poolCouponActiveParam = sync.Pool{
	New: func() any {
		return new(CouponActiveParam)
	},
}

// GetCouponActiveParam() 从对象池中获取CouponActiveParam
func GetCouponActiveParam() *CouponActiveParam {
	return poolCouponActiveParam.Get().(*CouponActiveParam)
}

// ReleaseCouponActiveParam 释放CouponActiveParam
func ReleaseCouponActiveParam(v *CouponActiveParam) {
	v.CouponEndDate = ""
	v.CouponInstanceId = ""
	v.CouponStartDate = ""
	poolCouponActiveParam.Put(v)
}
