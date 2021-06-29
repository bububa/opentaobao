package idle

// RentalOrderItemDto 
type RentalOrderItemDto struct {
    // 商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // sku信息
    SkuDTO   *ItemSkuDto `json:"sku_d_t_o,omitempty" xml:"sku_d_t_o,omitempty"`
    // 价格信息
    Price   *PriceDto `json:"price,omitempty" xml:"price,omitempty"`
}
