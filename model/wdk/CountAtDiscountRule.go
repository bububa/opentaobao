package wdk

// CountAtDiscountRule 
/* model for simplify = false
type CountAtDiscountRule struct {

    // 第N件折扣率【600=6折】
    
    CountAtDiscountRate   int64 `json:"count_at_discount_rate,omitempty"`
    

    // 第N件减钱【分】
    
    CountAtDecreaseMoney   int64 `json:"count_at_decrease_money,omitempty"`
    

    // 第N件一口价【分】
    
    CountAtFixPrice   int64 `json:"count_at_fix_price,omitempty"`
    

    // 是否第N建一口价
    
    IsCountAtFixPrice   bool `json:"is_count_at_fix_price,omitempty"`
    

    // 是否第N件减钱
    
    IsCountAtDecreaseMoney   bool `json:"is_count_at_decrease_money,omitempty"`
    

    // 是否第N件打折
    
    IsCountAtDiscountRate   bool `json:"is_count_at_discount_rate,omitempty"`
    

}
*/

// CountAtDiscountRule 
type CountAtDiscountRule struct {

    // 第N件折扣率【600=6折】
    CountAtDiscountRate   int64 `json:"count_at_discount_rate,omitempty"`

    // 第N件减钱【分】
    CountAtDecreaseMoney   int64 `json:"count_at_decrease_money,omitempty"`

    // 第N件一口价【分】
    CountAtFixPrice   int64 `json:"count_at_fix_price,omitempty"`

    // 是否第N建一口价
    IsCountAtFixPrice   bool `json:"is_count_at_fix_price,omitempty"`

    // 是否第N件减钱
    IsCountAtDecreaseMoney   bool `json:"is_count_at_decrease_money,omitempty"`

    // 是否第N件打折
    IsCountAtDiscountRate   bool `json:"is_count_at_discount_rate,omitempty"`

}
