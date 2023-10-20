package promotion

import (
	"sync"
)

// CouponTemplateInvestmentConfig 结构体
type CouponTemplateInvestmentConfig struct {
	// 出资人配置
	InvestmentInfoList []InvestmentInfo `json:"investment_info_list,omitempty" xml:"investment_info_list>investment_info,omitempty"`
}

var poolCouponTemplateInvestmentConfig = sync.Pool{
	New: func() any {
		return new(CouponTemplateInvestmentConfig)
	},
}

// GetCouponTemplateInvestmentConfig() 从对象池中获取CouponTemplateInvestmentConfig
func GetCouponTemplateInvestmentConfig() *CouponTemplateInvestmentConfig {
	return poolCouponTemplateInvestmentConfig.Get().(*CouponTemplateInvestmentConfig)
}

// ReleaseCouponTemplateInvestmentConfig 释放CouponTemplateInvestmentConfig
func ReleaseCouponTemplateInvestmentConfig(v *CouponTemplateInvestmentConfig) {
	v.InvestmentInfoList = v.InvestmentInfoList[:0]
	poolCouponTemplateInvestmentConfig.Put(v)
}
