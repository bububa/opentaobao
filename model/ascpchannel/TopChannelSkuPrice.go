package ascpchannel

// TopChannelSkuPrice 结构体
type TopChannelSkuPrice struct {
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku价格
	SkuPrice *SkuPrice `json:"sku_price,omitempty" xml:"sku_price,omitempty"`
}
