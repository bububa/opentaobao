package trade

// FastBuyPosCreateResult 
/* model for simplify = false
type FastBuyPosCreateResult struct {

    // 盒马订单号
    
    BizOrderId   int64 `json:"biz_order_id,omitempty"`
    

    // 优惠券优惠金额
    
    CouponFee   int64 `json:"coupon_fee,omitempty"`
    

    // 优惠活动优惠金额
    
    PromotionFee   int64 `json:"promotion_fee,omitempty"`
    

    // 返回错误码
    
    ReturnCode   string `json:"return_code,omitempty"`
    

    // 返回错误信息
    
    ReturnMsg   string `json:"return_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// FastBuyPosCreateResult 
type FastBuyPosCreateResult struct {

    // 盒马订单号
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 优惠券优惠金额
    CouponFee   int64 `json:"coupon_fee,omitempty"`

    // 优惠活动优惠金额
    PromotionFee   int64 `json:"promotion_fee,omitempty"`

    // 返回错误码
    ReturnCode   string `json:"return_code,omitempty"`

    // 返回错误信息
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
