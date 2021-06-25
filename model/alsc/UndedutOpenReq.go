package alsc

// UndedutOpenReq 
type UndedutOpenReq struct {

    // 品牌ID(不能和outbrandid同时为空)
    BrandId   string `json:"brand_id,omitempty"`

    // 卡ID，礼品卡或会员卡Id
    CardId   string `json:"card_id,omitempty"`

    // 会员ID
    CustomerId   string `json:"customer_id,omitempty"`

    // 外部反核销交易订单号(外部调用方保证在isv内部是唯一，可以是paymentItemId)
    NewOuterOrderId   string `json:"new_outer_order_id,omitempty"`

    // 操作人ID
    OperatorId   string `json:"operator_id,omitempty"`

    // 外部核销交易单号id(外部调用方保证在isv内部是唯一，可以是paymentItemId)
    OuterOrderId   string `json:"outer_order_id,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 请求id，用来做幂等
    RequestId   string `json:"request_id,omitempty"`

    // 门店ID(不能和outshopid同时为空)
    ShopId   string `json:"shop_id,omitempty"`

    // 当vlaue不传时，就是退整单金额
    Value   int64 `json:"value,omitempty"`

    // 外部门店ID(不能和shopid同时为空)
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 外部品牌ID(不能和brandid同时为空)
    OutBrandId   string `json:"out_brand_id,omitempty"`

    // CS是辰森，KRY是客如云
    BizChannel   string `json:"biz_channel,omitempty"`

}
