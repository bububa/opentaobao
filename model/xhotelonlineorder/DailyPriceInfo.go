package xhotelonlineorder

// DailyPriceInfo 
type DailyPriceInfo struct {

    // 日历早餐
    
    BreakFastCount   int64 `json:"break_fast_count,omitempty" xml:"break_fast_count,omitempty"`
    

    // 日历价格
    
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    

    // 日历日期
    
    Day   string `json:"day,omitempty" xml:"day,omitempty"`
    

    // 如果是低价加价商品，此价格是底价。如果是非底价商品且为会员商品，则为会员结算价
    
    BasePrice   int64 `json:"base_price,omitempty" xml:"base_price,omitempty"`
    

}
