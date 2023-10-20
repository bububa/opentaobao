package promotion

import (
	"sync"
)

// CouponTemplateParticipateConfig 结构体
type CouponTemplateParticipateConfig struct {
	// 参与者列表
	ParticipateList []LogicGroup `json:"participate_list,omitempty" xml:"participate_list>logic_group,omitempty"`
}

var poolCouponTemplateParticipateConfig = sync.Pool{
	New: func() any {
		return new(CouponTemplateParticipateConfig)
	},
}

// GetCouponTemplateParticipateConfig() 从对象池中获取CouponTemplateParticipateConfig
func GetCouponTemplateParticipateConfig() *CouponTemplateParticipateConfig {
	return poolCouponTemplateParticipateConfig.Get().(*CouponTemplateParticipateConfig)
}

// ReleaseCouponTemplateParticipateConfig 释放CouponTemplateParticipateConfig
func ReleaseCouponTemplateParticipateConfig(v *CouponTemplateParticipateConfig) {
	v.ParticipateList = v.ParticipateList[:0]
	poolCouponTemplateParticipateConfig.Put(v)
}
