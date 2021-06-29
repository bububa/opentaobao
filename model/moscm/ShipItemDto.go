package moscm

// ShipItemDTO 
type ShipItemDTO struct {
    // 数量
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 商品sku
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 唯一标识
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
}
