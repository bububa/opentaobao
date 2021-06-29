package aedropshiper

// PlaceOrderRequest4OpenApiDto 
type PlaceOrderRequest4OpenApiDto struct {

    // 物流地址信息
    
    LogisticsAddress   *MaillingAddressRequestDto `json:"logistics_address,omitempty" xml:"logistics_address,omitempty"`
    

    // 商品属性
    
    ProductItems   []ProductBaseItem `json:"product_items,omitempty" xml:"product_items,omitempty"`
    

}
