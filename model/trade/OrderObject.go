package trade

// OrderObject 
type OrderObject struct {

    // 优惠金额
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 用户昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 订单状态	PAID_DONE 已付款  	ACCEPT_ORDER 已接单  	PICK_ORDER 已拣货  	SHIPPED 已发货  	TRADE_SUCCESS 交易成功  	TRADE_CLOSE 交易关闭
    OrderStatus   string `json:"order_status,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 订单原金额
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 用户备注
    UserMem   string `json:"user_mem,omitempty"`

    // 订单履约状态	ACCEPTED 接单  	REJECTED 拒单  	CANCELED 取消订单  	SHIPPING 揽收（配送中）  	REFUSED 用户拒收  	SHIPPED  	TIMEOUTSHIPPED 超时签收  	RESHIPPED 二次妥投  	OUTOFSTOCK 缺货出  	FOODPROCESSDONE 加工完成  	PACKAGED 打包完成  	REFUNDED 逆向终态（已取货入站&退款完成）
    OrderFulfillStatus   string `json:"order_fulfill_status,omitempty"`

    // 配送人员信息
    Deliverer   *OrderDeliverer `json:"deliverer,omitempty"`

    // 实际支付金额
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 门店标识
    ShopId   string `json:"shop_id,omitempty"`

    // 买家标识
    UserId   string `json:"user_id,omitempty"`

    // 业务订单标识
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 收货人信息
    Delivery   *OrderDelivery `json:"delivery,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 子订单
    SubOrders   []Suborders `json:"sub_orders,omitempty"`

    // 业务订单标识
    OutOrderId   string `json:"out_order_id,omitempty"`

}
