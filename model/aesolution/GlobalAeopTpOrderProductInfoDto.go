package aesolution

// GlobalAeopTpOrderProductInfoDto 结构体
type GlobalAeopTpOrderProductInfoDto struct {
	// product SKU details
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// product name
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// Leaf category Id of the product
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// product quantity
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// product unit price
	UnitPrice *GlobalMoneyStr `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}
