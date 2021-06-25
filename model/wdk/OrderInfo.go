package wdk

// OrderInfo 
type OrderInfo struct {

    // 买家信息
    Buyer   *Buyer `json:"buyer,omitempty"`

    // 收件人信息
    Consignee   *Consignee `json:"consignee,omitempty"`

    // 子订单信息
    SubOrders   []SubOrder `json:"sub_orders,omitempty"`

    // 实际支付金额
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 原始金额
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 优惠金额
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 运费
    PostFee   int64 `json:"post_fee,omitempty"`

    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 订单状态
    OrderStatus   string `json:"order_status,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 创单时间
    CreateTime   string `json:"create_time,omitempty"`

    // 订单来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 外部渠道店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 配送方式 1:平台配送 2:商家自配送 3:到店自提
    PickupType   int64 `json:"pickup_type,omitempty"`

    // 平台佣金
    Commission   int64 `json:"commission,omitempty"`

    // 订单小号
    SerialNo   string `json:"serial_no,omitempty"`

    // 渠道店Id
    ShopId   string `json:"shop_id,omitempty"`

    // 经营店Id
    StoreId   string `json:"store_id,omitempty"`

}
