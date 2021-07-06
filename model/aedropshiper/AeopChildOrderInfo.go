package aedropshiper

// AeopChildOrderInfo 结构体
type AeopChildOrderInfo struct {
	// 商品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 商品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 商品价格
	ProductPrice *SimpleMoney `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 商品数量
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
}
