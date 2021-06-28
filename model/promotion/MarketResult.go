package promotion

// MarketResult 
type MarketResult struct {

    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 返回的匿名码对象
    
    Data   *CouponActivity `json:"data,omitempty" xml:"data,omitempty"`
    

}
