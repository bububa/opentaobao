package wdk

import (
	"sync"
)

// SeparatePricingDiscountRule 结构体
type SeparatePricingDiscountRule struct {
	// 是否为打折类型
	IsDiscountRate bool `json:"is_discount_rate,omitempty" xml:"is_discount_rate,omitempty"`
	// 是否为减钱类型
	IsDecrease bool `json:"is_decrease,omitempty" xml:"is_decrease,omitempty"`
	// 是否为一口价类型
	IsFixPrice bool `json:"is_fix_price,omitempty" xml:"is_fix_price,omitempty"`
}

var poolSeparatePricingDiscountRule = sync.Pool{
	New: func() any {
		return new(SeparatePricingDiscountRule)
	},
}

// GetSeparatePricingDiscountRule() 从对象池中获取SeparatePricingDiscountRule
func GetSeparatePricingDiscountRule() *SeparatePricingDiscountRule {
	return poolSeparatePricingDiscountRule.Get().(*SeparatePricingDiscountRule)
}

// ReleaseSeparatePricingDiscountRule 释放SeparatePricingDiscountRule
func ReleaseSeparatePricingDiscountRule(v *SeparatePricingDiscountRule) {
	v.IsDiscountRate = false
	v.IsDecrease = false
	v.IsFixPrice = false
	poolSeparatePricingDiscountRule.Put(v)
}
