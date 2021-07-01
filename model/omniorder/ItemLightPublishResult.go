package omniorder

// ItemLightPublishResult 结构体
type ItemLightPublishResult struct {
	// 创建生成的itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 创建生成的skuId
	SkuIds string `json:"sku_ids,omitempty" xml:"sku_ids,omitempty"`
	// 重复商品信息
	DuplicateInfos []ItemSkuDuplicateInfo `json:"duplicate_infos,omitempty" xml:"duplicate_infos>item_sku_duplicate_info,omitempty"`
}
