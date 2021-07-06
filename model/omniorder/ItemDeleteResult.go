package omniorder

// ItemDeleteResult 结构体
type ItemDeleteResult struct {
	// 重复商品信息
	DuplicateInfos []ItemSkuDuplicateInfo `json:"duplicate_infos,omitempty" xml:"duplicate_infos>item_sku_duplicate_info,omitempty"`
	// itemLightPublishDTO
	ItemLightPublishDTO *ItemLightPublishDto `json:"item_light_publish_d_t_o,omitempty" xml:"item_light_publish_d_t_o,omitempty"`
}
