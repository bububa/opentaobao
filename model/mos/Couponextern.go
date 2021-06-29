package mos

// Couponextern 
type Couponextern struct {
    // 券号
    CouponCode   string `json:"coupon_code,omitempty" xml:"coupon_code,omitempty"`
    // 券名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 券金额
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 券状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 券商品分摊
    CouponGoodApportions   []CouponGoodApportion `json:"coupon_good_apportions,omitempty" xml:"coupon_good_apportions>coupon_good_apportion,omitempty"`
}
