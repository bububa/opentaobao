package promotion

// CouponTemplateDiscountConfig 
/* model for simplify = false
type CouponTemplateDiscountConfig struct {

    // 是否减钱
    
    Decrease   bool `json:"decrease,omitempty"`
    

    // 减钱金额
    
    DecreaseMoney   int64 `json:"decrease_money,omitempty"`
    

    // 是否打折
    
    Discount   bool `json:"discount,omitempty"`
    

    // 打折额度
    
    DiscountRate   int64 `json:"discount_rate,omitempty"`
    

    // 优惠后的固定价格
    
    FixPriceAmount   int64 `json:"fix_price_amount,omitempty"`
    

    // 是否固定价格
    
    FixPrice   bool `json:"fix_price,omitempty"`
    

}
*/

// CouponTemplateDiscountConfig 
type CouponTemplateDiscountConfig struct {

    // 是否减钱
    Decrease   bool `json:"decrease,omitempty"`

    // 减钱金额
    DecreaseMoney   int64 `json:"decrease_money,omitempty"`

    // 是否打折
    Discount   bool `json:"discount,omitempty"`

    // 打折额度
    DiscountRate   int64 `json:"discount_rate,omitempty"`

    // 优惠后的固定价格
    FixPriceAmount   int64 `json:"fix_price_amount,omitempty"`

    // 是否固定价格
    FixPrice   bool `json:"fix_price,omitempty"`

}
