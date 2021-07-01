package ascpchannel

// Itemdolist 结构体
type Itemdolist struct {
	// 前端商品 ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 前端SKU ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
