package wdk

// SeparatePricingDiscountRule 
type SeparatePricingDiscountRule struct {
    // 是否为打折类型
    IsDiscountRate   bool `json:"is_discount_rate,omitempty" xml:"is_discount_rate,omitempty"`
    // 是否为减钱类型
    IsDecrease   bool `json:"is_decrease,omitempty" xml:"is_decrease,omitempty"`
    // 是否为一口价类型
    IsFixPrice   bool `json:"is_fix_price,omitempty" xml:"is_fix_price,omitempty"`
}
