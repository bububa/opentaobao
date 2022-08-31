package alihouse

// GoodsRelationDto 结构体
type GoodsRelationDto struct {
	// 交易商品所属商品id或者货所属商品id
	RelationItemId int64 `json:"relation_item_id,omitempty" xml:"relation_item_id,omitempty"`
	// 如果上翻的货为房源，则关联的为认购商品时，需必填此房源下认购商品所属的SKUid
	RelationSkuId int64 `json:"relation_sku_id,omitempty" xml:"relation_sku_id,omitempty"`
	// 类型 1-安心置业 2-特价房 3-0元购  4-大额电商券 5-认购商品 6-楼栋  7-户型  8-房源
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
