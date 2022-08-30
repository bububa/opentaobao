package ascp

// IcItemDto 结构体
type IcItemDto struct {
	// 店铺宝贝ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 店铺宝贝skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
