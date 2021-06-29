package trade

// TradeOrder 
type TradeOrder struct {
    // 门店标识
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
    // 业务订单标识
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 业务子订单标识(允许为null)
    SubBizOrderIds   []string `json:"sub_biz_order_ids,omitempty" xml:"sub_biz_order_ids>string,omitempty"`
    // 优惠金额
    DiscountFee   int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
    // 用户昵称
    UserNick   string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 订单原金额
    OriginFee   int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
    // 用户备注
    UserMem   string `json:"user_mem,omitempty" xml:"user_mem,omitempty"`
    // 支付金额(=整单原价-整单优惠金额+运费)
    PayFee   int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
    // 买家标识
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 收货人信息
    Delivery   *OrderDelivery `json:"delivery,omitempty" xml:"delivery,omitempty"`
    // 子订单
    SubOrders   []Suborders `json:"sub_orders,omitempty" xml:"sub_orders>suborders,omitempty"`
    // 外部关联标识
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 渠道业务, 如大润发优鲜=CHANNEL_HALFDAY
    BizType   string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 订单状态
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 订单履约状态
    OrderFulfillStatus   string `json:"order_fulfill_status,omitempty" xml:"order_fulfill_status,omitempty"`
    // 配送员信息
    Deliverer   *OrderDelivery `json:"deliverer,omitempty" xml:"deliverer,omitempty"`
    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
}
