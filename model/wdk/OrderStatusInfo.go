package wdk

// OrderStatusInfo 
/* model for simplify = false
type OrderStatusInfo struct {

    // 外部订单号
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 订单状态   已支付： PAID；  TRADE_CLOSE（仅未支付订单）
    
    OrderStatus   string `json:"order_status,omitempty"`
    

    // 外部渠道店ID
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 渠道来源
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

    // 经营店Id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 渠道店Id
    
    ShopId   string `json:"shop_id,omitempty"`
    

}
*/

// OrderStatusInfo 
type OrderStatusInfo struct {

    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 订单状态   已支付： PAID；  TRADE_CLOSE（仅未支付订单）
    OrderStatus   string `json:"order_status,omitempty"`

    // 外部渠道店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 经营店Id
    StoreId   string `json:"store_id,omitempty"`

    // 渠道店Id
    ShopId   string `json:"shop_id,omitempty"`

}
