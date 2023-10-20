package wdk

import (
	"sync"
)

// CouponStatisticsResultDo 结构体
type CouponStatisticsResultDo struct {
	// 券id
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 导购员id
	GuideId string `json:"guide_id,omitempty" xml:"guide_id,omitempty"`
	// 业务标识
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 日期
	StatisticsDate string `json:"statistics_date,omitempty" xml:"statistics_date,omitempty"`
	// 核券量
	UseCouponCount int64 `json:"use_coupon_count,omitempty" xml:"use_coupon_count,omitempty"`
	// 发券量
	SendCouponCount int64 `json:"send_coupon_count,omitempty" xml:"send_coupon_count,omitempty"`
}

var poolCouponStatisticsResultDo = sync.Pool{
	New: func() any {
		return new(CouponStatisticsResultDo)
	},
}

// GetCouponStatisticsResultDo() 从对象池中获取CouponStatisticsResultDo
func GetCouponStatisticsResultDo() *CouponStatisticsResultDo {
	return poolCouponStatisticsResultDo.Get().(*CouponStatisticsResultDo)
}

// ReleaseCouponStatisticsResultDo 释放CouponStatisticsResultDo
func ReleaseCouponStatisticsResultDo(v *CouponStatisticsResultDo) {
	v.CouponId = ""
	v.CouponName = ""
	v.GuideId = ""
	v.MerchantCode = ""
	v.StatisticsDate = ""
	v.UseCouponCount = 0
	v.SendCouponCount = 0
	poolCouponStatisticsResultDo.Put(v)
}
