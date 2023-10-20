package promotion

import (
	"sync"
)

// CouponTemplateTimeLimitConfig 结构体
type CouponTemplateTimeLimitConfig struct {
	// 优惠券结束时间
	EndValidTime string `json:"end_valid_time,omitempty" xml:"end_valid_time,omitempty"`
	// 优惠券开始时间
	StartValidTime string `json:"start_valid_time,omitempty" xml:"start_valid_time,omitempty"`
	// 优惠券有效时间类型 RANGE(1,&#34;开始/结束时间&#34;), DURATION(2,&#34;固定时长&#34;),BETWEEN(4, &#34;发放后X-Y天失效&#34;)
	ValidTimeType int64 `json:"valid_time_type,omitempty" xml:"valid_time_type,omitempty"`
	// 优惠券有效时长，单位为秒（固定有效时长的优惠券）
	ValidityPeriod int64 `json:"validity_period,omitempty" xml:"validity_period,omitempty"`
	// 优惠券有效开始时长，单位为秒（优惠券领取后X天开始有效）
	ValidityStartInterval int64 `json:"validity_start_interval,omitempty" xml:"validity_start_interval,omitempty"`
	// 优惠券有效结束时长，单位为秒（优惠券领取后X-Y天有效的Y）
	ValidityEndInterval int64 `json:"validity_end_interval,omitempty" xml:"validity_end_interval,omitempty"`
}

var poolCouponTemplateTimeLimitConfig = sync.Pool{
	New: func() any {
		return new(CouponTemplateTimeLimitConfig)
	},
}

// GetCouponTemplateTimeLimitConfig() 从对象池中获取CouponTemplateTimeLimitConfig
func GetCouponTemplateTimeLimitConfig() *CouponTemplateTimeLimitConfig {
	return poolCouponTemplateTimeLimitConfig.Get().(*CouponTemplateTimeLimitConfig)
}

// ReleaseCouponTemplateTimeLimitConfig 释放CouponTemplateTimeLimitConfig
func ReleaseCouponTemplateTimeLimitConfig(v *CouponTemplateTimeLimitConfig) {
	v.EndValidTime = ""
	v.StartValidTime = ""
	v.ValidTimeType = 0
	v.ValidityPeriod = 0
	v.ValidityStartInterval = 0
	v.ValidityEndInterval = 0
	poolCouponTemplateTimeLimitConfig.Put(v)
}
