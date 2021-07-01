package icbu

// BulkDiscountPrice 结构体
type BulkDiscountPrice struct {
	// 起始数量，范围是1-99999
	StartQuantity int64 `json:"start_quantity,omitempty" xml:"start_quantity,omitempty"`
	// 价格，范围是0.01-9999999.00
	Price string `json:"price,omitempty" xml:"price,omitempty"`
}
