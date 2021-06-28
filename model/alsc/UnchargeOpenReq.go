package alsc

// UnchargeOpenReq 
/* model for simplify = false
type UnchargeOpenReq struct {

    // SaaS品牌ID(不能和outbrandid同时为空)
    
    BrandId   string `json:"brand_id,omitempty"`
    

    // 外部品牌ID(不能和brandid同时为空)
    
    OutBrandId   string `json:"out_brand_id,omitempty"`
    

    // SaaS门店ID(不能和outshopid同时为空)
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 外部门店ID(不能和sopid同时为空)
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 卡号
    
    CardId   string `json:"card_id,omitempty"`
    

    // 操作人ID(SaaS Id）
    
    OperatorId   string `json:"operator_id,omitempty"`
    

    // 退储值订单Id，必填（kry是paymentItemId）,外部保证唯一，作为退款幂等号
    
    OuterOrderId   string `json:"outer_order_id,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 请求id，幂等
    
    RequestId   string `json:"request_id,omitempty"`
    

    // 原充值订单Id，必填（kry是srcPaymentItemId）
    
    SrcOuterOrderId   string `json:"src_outer_order_id,omitempty"`
    

    // CS是辰森，KRY是客如云
    
    BizChannel   string `json:"biz_channel,omitempty"`
    

    // 原支付单号
    
    SrcOutPayId   string `json:"src_out_pay_id,omitempty"`
    

}
*/

// UnchargeOpenReq 
type UnchargeOpenReq struct {

    // SaaS品牌ID(不能和outbrandid同时为空)
    BrandId   string `json:"brand_id,omitempty"`

    // 外部品牌ID(不能和brandid同时为空)
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // SaaS门店ID(不能和outshopid同时为空)
    ShopId   string `json:"shop_id,omitempty"`

    // 外部门店ID(不能和sopid同时为空)
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 卡号
    CardId   string `json:"card_id,omitempty"`

    // 操作人ID(SaaS Id）
    OperatorId   string `json:"operator_id,omitempty"`

    // 退储值订单Id，必填（kry是paymentItemId）,外部保证唯一，作为退款幂等号
    OuterOrderId   string `json:"outer_order_id,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 请求id，幂等
    RequestId   string `json:"request_id,omitempty"`

    // 原充值订单Id，必填（kry是srcPaymentItemId）
    SrcOuterOrderId   string `json:"src_outer_order_id,omitempty"`

    // CS是辰森，KRY是客如云
    BizChannel   string `json:"biz_channel,omitempty"`

    // 原支付单号
    SrcOutPayId   string `json:"src_out_pay_id,omitempty"`

}
