package inventory

// RelationList 
type RelationList struct {
    // 生效的前端宝贝id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 生效的前端宝贝的skuid
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
