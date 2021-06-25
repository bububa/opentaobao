package trade

// OrderDiscountBillBo 
type OrderDiscountBillBo struct {

    // 活动ID
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 活动名称
    ActivityName   string `json:"activity_name,omitempty"`

    // 活动类型 1：活动 2：券
    ActivityType   int64 `json:"activity_type,omitempty"`

    // 购买数量
    BuyQuantity   int64 `json:"buy_quantity,omitempty"`

    // 优惠金额
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 优惠件数
    DiscountQuantity   int64 `json:"discount_quantity,omitempty"`

    // 补差类型
    DiscountType   int64 `json:"discount_type,omitempty"`

    // 主键id
    Id   int64 `json:"id,omitempty"`

    // 主订单号
    MainOrderId   string `json:"main_order_id,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 商家优惠补差金额
    MerchantDiscountFee   int64 `json:"merchant_discount_fee,omitempty"`

    // 交易状态
    OrderStatus   int64 `json:"order_status,omitempty"`

    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

    // 子单号
    SubOrderId   string `json:"sub_order_id,omitempty"`

    // 淘鲜达优惠金额
    TxdDiscountFee   int64 `json:"txd_discount_fee,omitempty"`

    // 业务时间
    BizTime   string `json:"biz_time,omitempty"`

    // 订单渠道
    OrderChannel   int64 `json:"order_channel,omitempty"`

}
