package promotion

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
