package wdk

// InboundItemInfo 
type InboundItemInfo struct {
    // 收货数量
    InboundQuantity   string `json:"inbound_quantity,omitempty" xml:"inbound_quantity,omitempty"`
    // 商品编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}
