package trade

// FastBuyPosQueryResult 
type FastBuyPosQueryResult struct {

    // 五道口订单id
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 优惠券优惠金额
    CouponFee   int64 `json:"coupon_fee,omitempty"`

    // 订单状态：1为已下单，未支付；2为交易完成；3为已退款；4为已关单
    OrderStatus   int64 `json:"order_status,omitempty"`

    // 优惠活动优惠金额
    PromotionFee   int64 `json:"promotion_fee,omitempty"`

    // 返回的错误码
    ReturnCode   string `json:"return_code,omitempty"`

    // 错误信息
    ReturnMsg   string `json:"return_msg,omitempty"`

    // 是否调用成功
    Success   bool `json:"success,omitempty"`

    // 商品分摊优惠
    ItemPromotions   string `json:"item_promotions,omitempty"`

}
