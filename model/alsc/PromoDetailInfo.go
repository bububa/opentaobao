package alsc

// PromoDetailInfo 
/* model for simplify = false
type PromoDetailInfo struct {

    // 优惠授权人id
    
    OutAuthorizerId   string `json:"out_authorizer_id,omitempty"`
    

    // 业务授权人名称
    
    OutAuthorizerName   string `json:"out_authorizer_name,omitempty"`
    

    // 关联订单号
    
    OutOrderNo   string `json:"out_order_no,omitempty"`
    

    // 优惠明细ID
    
    OutPromoDetailId   string `json:"out_promo_detail_id,omitempty"`
    

    // 优惠金额，单位“分”，减钱为负值
    
    PromoFee   int64 `json:"promo_fee,omitempty"`
    

    // 优惠id，券id or 活动id or other等
    
    PromoId   string `json:"promo_id,omitempty"`
    

    // 优惠名称
    
    PromoName   string `json:"promo_name,omitempty"`
    

    // 优惠原因：COUPON券优惠  PROMO促销优惠  SALES特价商品优惠 POINTS_DEDUCT积分抵现  SALES_MEMBER会员价商品优惠 MODIFY_PRICE手动改价 OTHER其他
    
    PromoReason   string `json:"promo_reason,omitempty"`
    

}
*/

// PromoDetailInfo 
type PromoDetailInfo struct {

    // 优惠授权人id
    OutAuthorizerId   string `json:"out_authorizer_id,omitempty"`

    // 业务授权人名称
    OutAuthorizerName   string `json:"out_authorizer_name,omitempty"`

    // 关联订单号
    OutOrderNo   string `json:"out_order_no,omitempty"`

    // 优惠明细ID
    OutPromoDetailId   string `json:"out_promo_detail_id,omitempty"`

    // 优惠金额，单位“分”，减钱为负值
    PromoFee   int64 `json:"promo_fee,omitempty"`

    // 优惠id，券id or 活动id or other等
    PromoId   string `json:"promo_id,omitempty"`

    // 优惠名称
    PromoName   string `json:"promo_name,omitempty"`

    // 优惠原因：COUPON券优惠  PROMO促销优惠  SALES特价商品优惠 POINTS_DEDUCT积分抵现  SALES_MEMBER会员价商品优惠 MODIFY_PRICE手动改价 OTHER其他
    PromoReason   string `json:"promo_reason,omitempty"`

}
