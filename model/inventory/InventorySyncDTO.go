package inventory

// InventorySyncDTO 
type InventorySyncDTO struct {
    // 商品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 库存数量
    ItemAmount   int64 `json:"item_amount,omitempty" xml:"item_amount,omitempty"`
}
