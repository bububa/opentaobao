package promotion

import (
	"sync"
)

// CouponTemplateDiscountConfig 结构体
type CouponTemplateDiscountConfig struct {
	// 减钱金额
	DecreaseMoney int64 `json:"decrease_money,omitempty" xml:"decrease_money,omitempty"`
	// 打折额度
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 优惠后的固定价格
	FixPriceAmount int64 `json:"fix_price_amount,omitempty" xml:"fix_price_amount,omitempty"`
	// 是否减钱
	Decrease bool `json:"decrease,omitempty" xml:"decrease,omitempty"`
	// 是否打折
	Discount bool `json:"discount,omitempty" xml:"discount,omitempty"`
	// 是否固定价格
	FixPrice bool `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
}

var poolCouponTemplateDiscountConfig = sync.Pool{
	New: func() any {
		return new(CouponTemplateDiscountConfig)
	},
}

// GetCouponTemplateDiscountConfig() 从对象池中获取CouponTemplateDiscountConfig
func GetCouponTemplateDiscountConfig() *CouponTemplateDiscountConfig {
	return poolCouponTemplateDiscountConfig.Get().(*CouponTemplateDiscountConfig)
}

// ReleaseCouponTemplateDiscountConfig 释放CouponTemplateDiscountConfig
func ReleaseCouponTemplateDiscountConfig(v *CouponTemplateDiscountConfig) {
	v.DecreaseMoney = 0
	v.DiscountRate = 0
	v.FixPriceAmount = 0
	v.Decrease = false
	v.Discount = false
	v.FixPrice = false
	poolCouponTemplateDiscountConfig.Put(v)
}
