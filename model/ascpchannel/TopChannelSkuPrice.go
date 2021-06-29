package ascpchannel

// TopChannelSkuPrice 
type TopChannelSkuPrice struct {
    // sku价格
    SkuPrice   *SkuPrice `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
    // skuId
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
