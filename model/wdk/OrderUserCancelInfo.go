package wdk

// OrderUserCancelInfo 
/* model for simplify = false
type OrderUserCancelInfo struct {

    // 外部订单ID
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 外部渠道店ID
    
    OutShopId   string `json:"out_shop_id,omitempty"`
    

    // 渠道来源
    
    OrderFrom   int64 `json:"order_from,omitempty"`
    

    // 盒马主单号
    
    BizOrderId   string `json:"biz_order_id,omitempty"`
    

    // 经营店Id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 渠道店Id
    
    ShopId   string `json:"shop_id,omitempty"`
    

}
*/

// OrderUserCancelInfo 
type OrderUserCancelInfo struct {

    // 外部订单ID
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 外部渠道店ID
    OutShopId   string `json:"out_shop_id,omitempty"`

    // 渠道来源
    OrderFrom   int64 `json:"order_from,omitempty"`

    // 盒马主单号
    BizOrderId   string `json:"biz_order_id,omitempty"`

    // 经营店Id
    StoreId   string `json:"store_id,omitempty"`

    // 渠道店Id
    ShopId   string `json:"shop_id,omitempty"`

}
