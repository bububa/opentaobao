package wdk

import (
	"sync"
)

// CountAtDiscountRule 结构体
type CountAtDiscountRule struct {
	// 第N件折扣率【600=6折】
	CountAtDiscountRate int64 `json:"count_at_discount_rate,omitempty" xml:"count_at_discount_rate,omitempty"`
	// 第N件减钱【分】
	CountAtDecreaseMoney int64 `json:"count_at_decrease_money,omitempty" xml:"count_at_decrease_money,omitempty"`
	// 第N件一口价【分】
	CountAtFixPrice int64 `json:"count_at_fix_price,omitempty" xml:"count_at_fix_price,omitempty"`
	// 是否第N建一口价
	IsCountAtFixPrice bool `json:"is_count_at_fix_price,omitempty" xml:"is_count_at_fix_price,omitempty"`
	// 是否第N件减钱
	IsCountAtDecreaseMoney bool `json:"is_count_at_decrease_money,omitempty" xml:"is_count_at_decrease_money,omitempty"`
	// 是否第N件打折
	IsCountAtDiscountRate bool `json:"is_count_at_discount_rate,omitempty" xml:"is_count_at_discount_rate,omitempty"`
}

var poolCountAtDiscountRule = sync.Pool{
	New: func() any {
		return new(CountAtDiscountRule)
	},
}

// GetCountAtDiscountRule() 从对象池中获取CountAtDiscountRule
func GetCountAtDiscountRule() *CountAtDiscountRule {
	return poolCountAtDiscountRule.Get().(*CountAtDiscountRule)
}

// ReleaseCountAtDiscountRule 释放CountAtDiscountRule
func ReleaseCountAtDiscountRule(v *CountAtDiscountRule) {
	v.CountAtDiscountRate = 0
	v.CountAtDecreaseMoney = 0
	v.CountAtFixPrice = 0
	v.IsCountAtFixPrice = false
	v.IsCountAtDecreaseMoney = false
	v.IsCountAtDiscountRate = false
	poolCountAtDiscountRule.Put(v)
}
