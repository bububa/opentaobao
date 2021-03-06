package aedropshiper

// PlaceOrderRequest4OpenApiDto 结构体
type PlaceOrderRequest4OpenApiDto struct {
	// 商品属性
	ProductItems []ProductBaseItem `json:"product_items,omitempty" xml:"product_items>product_base_item,omitempty"`
	// 物流地址信息
	LogisticsAddress *MaillingAddressRequestDto `json:"logistics_address,omitempty" xml:"logistics_address,omitempty"`
}
