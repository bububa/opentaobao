package promotion

import (
	"sync"
)

// CouponTemplateUserCrowdConfig 结构体
type CouponTemplateUserCrowdConfig struct {
	// 商家自定义人群
	MerchantCustomizeCrowdName string `json:"merchant_customize_crowd_name,omitempty" xml:"merchant_customize_crowd_name,omitempty"`
	// 平台人群
	PlatformCrowdName string `json:"platform_crowd_name,omitempty" xml:"platform_crowd_name,omitempty"`
}

var poolCouponTemplateUserCrowdConfig = sync.Pool{
	New: func() any {
		return new(CouponTemplateUserCrowdConfig)
	},
}

// GetCouponTemplateUserCrowdConfig() 从对象池中获取CouponTemplateUserCrowdConfig
func GetCouponTemplateUserCrowdConfig() *CouponTemplateUserCrowdConfig {
	return poolCouponTemplateUserCrowdConfig.Get().(*CouponTemplateUserCrowdConfig)
}

// ReleaseCouponTemplateUserCrowdConfig 释放CouponTemplateUserCrowdConfig
func ReleaseCouponTemplateUserCrowdConfig(v *CouponTemplateUserCrowdConfig) {
	v.MerchantCustomizeCrowdName = ""
	v.PlatformCrowdName = ""
	poolCouponTemplateUserCrowdConfig.Put(v)
}
