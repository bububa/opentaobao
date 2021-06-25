package wdk

// WdkOrderSyncBo 
type WdkOrderSyncBo struct {

    // 五道口订单ID
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 外部订单ID
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 渠道店ID
    ShopId   string `json:"shop_id,omitempty"`

    // 商户编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 订单来源, 如 TAOBAO (4, "TC自营渠道"),
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 买家open_uid
    OpenUid   string `json:"open_uid,omitempty"`

    // 用户订单支付金额,分
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 订单原价,分
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 订单优惠金额,分
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 订单配送费,分
    PostFee   int64 `json:"post_fee,omitempty"`

    // 订单状态,如PAID_DONE(2, "已付款"), TRADE_SUCCESS(6, "交易成功")
    OrderStatus   int64 `json:"order_status,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 期望送达时间段
    ExpectArriveTime   string `json:"expect_arrive_time,omitempty"`

    // 配送方式, 如InTime(1, "即时达"), SetTime(2, "定时达"),TopSpeed(3,"极速达"),NoNeedSend(4,"无需配送")
    ArriveType   int64 `json:"arrive_type,omitempty"`

    // 子单列表
    SubOrders   []Suborders `json:"sub_orders,omitempty"`

    // 订单优惠信息
    Promotions   []Promotions `json:"promotions,omitempty"`

}
