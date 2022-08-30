package aedropshiper

// AeChildOrderInfo 结构体
type AeChildOrderInfo struct {
	// Item name
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// Item ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// Item price
	ProductPrice *SimpleMoney `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// Item quantity
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
}
