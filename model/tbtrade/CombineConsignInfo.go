package tbtrade

// CombineConsignInfo 结构体
type CombineConsignInfo struct {
	// 套餐组合id
	CombId string `json:"comb_id,omitempty" xml:"comb_id,omitempty"`
	// 套餐内的商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 套餐内商品的skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 成分品的预计发货时间
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 成分品的预计发货时间
	RenderConsignTime string `json:"render_consign_time,omitempty" xml:"render_consign_time,omitempty"`
}
