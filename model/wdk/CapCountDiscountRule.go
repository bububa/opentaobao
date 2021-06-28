package wdk

// CapCountDiscountRule 
/* model for simplify = false
type CapCountDiscountRule struct {

    // 指定件数每件一口价【分】
    
    CapCountEachFixPrice   int64 `json:"cap_count_each_fix_price,omitempty"`
    

    // 指定件数整体打几折【600=6折】
    
    CapCountDiscountRate   int64 `json:"cap_count_discount_rate,omitempty"`
    

    // 指定件数整体减钱【分】
    
    CapCountDecreaseMoney   int64 `json:"cap_count_decrease_money,omitempty"`
    

    // 指定件数整体一口价【分】
    
    CapCountFixPrice   int64 `json:"cap_count_fix_price,omitempty"`
    

    // 是否指定件数整体一口价
    
    IsCapCountFixPrice   bool `json:"is_cap_count_fix_price,omitempty"`
    

    // 是否指定件数整体减钱
    
    IsCapCountDecreaseMoney   bool `json:"is_cap_count_decrease_money,omitempty"`
    

    // 是否指定件数整体打折
    
    IsCapCountDiscountRate   bool `json:"is_cap_count_discount_rate,omitempty"`
    

    // 是否指定件数每件一口价
    
    IsCapCountEachFixPrice   bool `json:"is_cap_count_each_fix_price,omitempty"`
    

}
*/

// CapCountDiscountRule 
type CapCountDiscountRule struct {

    // 指定件数每件一口价【分】
    CapCountEachFixPrice   int64 `json:"cap_count_each_fix_price,omitempty"`

    // 指定件数整体打几折【600=6折】
    CapCountDiscountRate   int64 `json:"cap_count_discount_rate,omitempty"`

    // 指定件数整体减钱【分】
    CapCountDecreaseMoney   int64 `json:"cap_count_decrease_money,omitempty"`

    // 指定件数整体一口价【分】
    CapCountFixPrice   int64 `json:"cap_count_fix_price,omitempty"`

    // 是否指定件数整体一口价
    IsCapCountFixPrice   bool `json:"is_cap_count_fix_price,omitempty"`

    // 是否指定件数整体减钱
    IsCapCountDecreaseMoney   bool `json:"is_cap_count_decrease_money,omitempty"`

    // 是否指定件数整体打折
    IsCapCountDiscountRate   bool `json:"is_cap_count_discount_rate,omitempty"`

    // 是否指定件数每件一口价
    IsCapCountEachFixPrice   bool `json:"is_cap_count_each_fix_price,omitempty"`

}
