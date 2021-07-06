package tttm

// ItemInfoDto 结构体
type ItemInfoDto struct {
	// 货品信息
	ProductInfoDTOs []ProductInfoDto `json:"product_info_d_t_os,omitempty" xml:"product_info_d_t_os>product_info_dto,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
