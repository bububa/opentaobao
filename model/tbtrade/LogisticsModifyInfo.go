package tbtrade

// LogisticsModifyInfo 结构体
type LogisticsModifyInfo struct {
	// 修改关联的订单好
	RelatedId string `json:"related_id,omitempty" xml:"related_id,omitempty"`
	// 修改后的发货时效
	ConsignTime string `json:"consign_time,omitempty" xml:"consign_time,omitempty"`
	// 修改前的发货时效
	OriginConsignTime string `json:"origin_consign_time,omitempty" xml:"origin_consign_time,omitempty"`
	// 修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 修改原因
	ModifyReason string `json:"modify_reason,omitempty" xml:"modify_reason,omitempty"`
	// 成分品的商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 成分品的skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 成分品的组合id
	CombId string `json:"comb_id,omitempty" xml:"comb_id,omitempty"`
}
