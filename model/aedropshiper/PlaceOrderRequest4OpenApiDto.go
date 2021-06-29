package aedropshiper

// PlaceOrderRequest4OpenApiDTO 
type PlaceOrderRequest4OpenApiDTO struct {
    // 物流地址信息
    LogisticsAddress   *MaillingAddressRequestDTO `json:"logistics_address,omitempty" xml:"logistics_address,omitempty"`
    // 商品属性
    ProductItems   []ProductBaseItem `json:"product_items,omitempty" xml:"product_items>product_base_item,omitempty"`
}
