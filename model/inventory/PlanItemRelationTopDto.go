package inventory

// PlanItemRelationTopDTO 
type PlanItemRelationTopDTO struct {
    // 计划生效的itemid
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 计划生效的skuid
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
