package aedropshiper

// ProductBaseItem 
type ProductBaseItem struct {
    // 商品数量
    ProductCount   int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
    // 商品id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 商品sku
    SkuAttr   string `json:"sku_attr,omitempty" xml:"sku_attr,omitempty"`
    // 物流服务名称
    LogisticsServiceName   string `json:"logistics_service_name,omitempty" xml:"logistics_service_name,omitempty"`
    // 用户留言
    OrderMemo   string `json:"order_memo,omitempty" xml:"order_memo,omitempty"`
}
