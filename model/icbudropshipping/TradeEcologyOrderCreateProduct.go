package icbudropshipping

// TradeEcologyOrderCreateProduct 
type TradeEcologyOrderCreateProduct struct {
    // product id
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // quantity
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // sku id
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // unit price
    UnitPriceStr   string `json:"unit_price_str,omitempty" xml:"unit_price_str,omitempty"`
}
