package mos

// Couponextern 
/* model for simplify = false
type Couponextern struct {

    // 券号
    
    CouponCode   string `json:"coupon_code,omitempty"`
    

    // 券名
    
    Name   string `json:"name,omitempty"`
    

    // 券金额
    
    Amount   string `json:"amount,omitempty"`
    

    // 券状态
    
    Status   string `json:"status,omitempty"`
    

    // 券商品分摊
    
    CouponGoodApportions  struct {
        CouponGoodApportion  []CouponGoodApportion `json:"coupon_good_apportion,omitempty"`
    } `json:"coupon_good_apportions,omitempty"`
    

}
*/

// Couponextern 
type Couponextern struct {

    // 券号
    CouponCode   string `json:"coupon_code,omitempty"`

    // 券名
    Name   string `json:"name,omitempty"`

    // 券金额
    Amount   string `json:"amount,omitempty"`

    // 券状态
    Status   string `json:"status,omitempty"`

    // 券商品分摊
    CouponGoodApportions   []CouponGoodApportion `json:"coupon_good_apportions,omitempty"`

}
