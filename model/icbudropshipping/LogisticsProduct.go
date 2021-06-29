package icbudropshipping

// LogisticsProduct 
type LogisticsProduct struct {
    // Product Id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // quantity
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // Sku ID
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
