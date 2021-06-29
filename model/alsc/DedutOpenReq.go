package alsc

// DedutOpenReq 
type DedutOpenReq struct {
    // 时间
    BizDate   string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
    // SaaS品牌ID(不能和outbrandid同时为空)
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 卡Id，礼品卡或会员卡Id
    CardId   string `json:"card_id,omitempty" xml:"card_id,omitempty"`
    // 会员
    CustomerId   string `json:"customer_id,omitempty" xml:"customer_id,omitempty"`
    // 操作人ID(SaaS Id）
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    // 外部核销交易单号id(外部调用方保证在isv内部是唯一，可以是paymentItemId)
    OuterOrderId   string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 请求id，用来做幂等
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // SaaS门店ID(不能和outshopid同时为空)
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 核销总资产
    Value   int64 `json:"value,omitempty" xml:"value,omitempty"`
    // 外部品牌ID(不能和brandid同时为空)
    OutBrandId   string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
    // CS是辰森，KRY是客如云
    BizChannel   string `json:"biz_channel,omitempty" xml:"biz_channel,omitempty"`
    // 外部门店ID,和门店ID必传一
    OutShopId   string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
    // 自定义参数按照json格式传入
    ExtInfo   string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}
