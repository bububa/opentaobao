package wdk

// InboundItemInfo 
/* model for simplify = false
type InboundItemInfo struct {

    // 收货数量
    
    InboundQuantity   string `json:"inbound_quantity,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

}
*/

// InboundItemInfo 
type InboundItemInfo struct {

    // 收货数量
    InboundQuantity   string `json:"inbound_quantity,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

}
