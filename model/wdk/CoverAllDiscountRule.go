package wdk

import (
	"sync"
)

// CoverAllDiscountRule 结构体
type CoverAllDiscountRule struct {
	// 每件商品一口价【分】
	EachFixPrice int64 `json:"each_fix_price,omitempty" xml:"each_fix_price,omitempty"`
	// 整体打折【600=6折】
	CoverAllDiscountRate int64 `json:"cover_all_discount_rate,omitempty" xml:"cover_all_discount_rate,omitempty"`
	// 整体减钱【分】
	CoverAllDecreaseMoney int64 `json:"cover_all_decrease_money,omitempty" xml:"cover_all_decrease_money,omitempty"`
	// 整体一口价【分】
	CoverAllFixPrice int64 `json:"cover_all_fix_price,omitempty" xml:"cover_all_fix_price,omitempty"`
	// 是否整体一口价
	IsCoverAllFixPrice bool `json:"is_cover_all_fix_price,omitempty" xml:"is_cover_all_fix_price,omitempty"`
	// 是否整体减钱
	IsCoverAllDecreaseMoney bool `json:"is_cover_all_decrease_money,omitempty" xml:"is_cover_all_decrease_money,omitempty"`
	// 是否整体打折
	IsCoverAllDiscountRate bool `json:"is_cover_all_discount_rate,omitempty" xml:"is_cover_all_discount_rate,omitempty"`
	// 是否每件一口价
	IsEachFixPrice bool `json:"is_each_fix_price,omitempty" xml:"is_each_fix_price,omitempty"`
}

var poolCoverAllDiscountRule = sync.Pool{
	New: func() any {
		return new(CoverAllDiscountRule)
	},
}

// GetCoverAllDiscountRule() 从对象池中获取CoverAllDiscountRule
func GetCoverAllDiscountRule() *CoverAllDiscountRule {
	return poolCoverAllDiscountRule.Get().(*CoverAllDiscountRule)
}

// ReleaseCoverAllDiscountRule 释放CoverAllDiscountRule
func ReleaseCoverAllDiscountRule(v *CoverAllDiscountRule) {
	v.EachFixPrice = 0
	v.CoverAllDiscountRate = 0
	v.CoverAllDecreaseMoney = 0
	v.CoverAllFixPrice = 0
	v.IsCoverAllFixPrice = false
	v.IsCoverAllDecreaseMoney = false
	v.IsCoverAllDiscountRate = false
	v.IsEachFixPrice = false
	poolCoverAllDiscountRule.Put(v)
}
