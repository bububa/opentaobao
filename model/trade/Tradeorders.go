package trade

// Tradeorders 
/* model for simplify = false
type Tradeorders struct {

    // 折扣金额
    
    DiscountFee   int64 `json:"discount_fee,omitempty"`
    

    // 买家昵称
    
    UserNick   string `json:"user_nick,omitempty"`
    

    // 订单状态
    
    OrderStatus   string `json:"order_status,omitempty"`
    

    // 支付时间
    
    PayTime   string `json:"pay_time,omitempty"`
    

    // 原金额
    
    OriginFee   int64 `json:"origin_fee,omitempty"`
    

    // 买家留言
    
    UserMem   string `json:"user_mem,omitempty"`
    

    // 订单履约状态
    
    OrderFulfillStatus   string `json:"order_fulfill_status,omitempty"`
    

    // 订单配送信息
    
    Deliverer  *struct {
        OrderDeliverer  *OrderDeliverer `json:"order_deliverer,omitempty"`
    } `json:"deliverer,omitempty"`
    

    // 支付金额
    
    PayFee   int64 `json:"pay_fee,omitempty"`
    

    // 门店编码
    
    ShopId   string `json:"shop_id,omitempty"`
    

    // 买家标识
    
    UserId   string `json:"user_id,omitempty"`
    

    // 业务订单标识
    
    BizOrderId   string `json:"biz_order_id,omitempty"`
    

    // 收货人信息
    
    Delivery  *struct {
        OrderDelivery  *OrderDelivery `json:"order_delivery,omitempty"`
    } `json:"delivery,omitempty"`
    

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 子订单
    
    SubOrders  struct {
        Suborders  []Suborders `json:"suborders,omitempty"`
    } `json:"sub_orders,omitempty"`
    

    // 外部关联订单标识
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

}
*/

// Tradeorders 
type Tradeorders struct {

    // 折扣金额
    DiscountFee   int64 `json:"discount_fee,omitempty"`

    // 买家昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 订单状态
    OrderStatus   string `json:"order_status,omitempty"`

    // 支付时间
    PayTime   string `json:"pay_time,omitempty"`

    // 原金额
    OriginFee   int64 `json:"origin_fee,omitempty"`

    // 买家留言
    UserMem   string `json:"user_mem,omitempty"`

    // 订单履约状态
    OrderFulfillStatus   string `json:"order_fulfill_status,omitempty"`

    // 订单配送信息
    Deliverer   *OrderDeliverer `json:"deliverer,omitempty"`

    // 支付金额
    PayFee   int64 `json:"pay_fee,omitempty"`

    // 门店编码
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

    // 外部关联订单标识
    OutOrderId   string `json:"out_order_id,omitempty"`

}
