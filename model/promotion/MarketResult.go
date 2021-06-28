package promotion

// MarketResult 
/* model for simplify = false
type MarketResult struct {

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 返回的匿名码对象
    
    Data  *struct {
        CouponActivity  *CouponActivity `json:"coupon_activity,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

// MarketResult 
type MarketResult struct {

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 返回的匿名码对象
    Data   *CouponActivity `json:"data,omitempty"`

}
