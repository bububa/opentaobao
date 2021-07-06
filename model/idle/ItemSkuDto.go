package idle

// ItemSkuDto 结构体
type ItemSkuDto struct {
	// sku属性
	PropList []ItemPvPairDto `json:"prop_list,omitempty" xml:"prop_list>item_pv_pair_dto,omitempty"`
	// sku价格，单位分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// sku库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
