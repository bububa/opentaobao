package promotion

import (
	"sync"
)

// CouponTemplateCompatibleConfig 结构体
type CouponTemplateCompatibleConfig struct {
	// 优惠券应用类型 pointCoupon：积分券
	ApplicationType string `json:"application_type,omitempty" xml:"application_type,omitempty"`
	// 是否要求优惠券在一天的23:59:59失效 1表示最后一秒失效
	ValidTillNight int64 `json:"valid_till_night,omitempty" xml:"valid_till_night,omitempty"`
}

var poolCouponTemplateCompatibleConfig = sync.Pool{
	New: func() any {
		return new(CouponTemplateCompatibleConfig)
	},
}

// GetCouponTemplateCompatibleConfig() 从对象池中获取CouponTemplateCompatibleConfig
func GetCouponTemplateCompatibleConfig() *CouponTemplateCompatibleConfig {
	return poolCouponTemplateCompatibleConfig.Get().(*CouponTemplateCompatibleConfig)
}

// ReleaseCouponTemplateCompatibleConfig 释放CouponTemplateCompatibleConfig
func ReleaseCouponTemplateCompatibleConfig(v *CouponTemplateCompatibleConfig) {
	v.ApplicationType = ""
	v.ValidTillNight = 0
	poolCouponTemplateCompatibleConfig.Put(v)
}
