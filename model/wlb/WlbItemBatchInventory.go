package wlb

// WlbItemBatchInventory 结构体
type WlbItemBatchInventory struct {
	// 批次库存查询结果
	ItemBatch []WlbItemBatch `json:"item_batch,omitempty" xml:"item_batch>wlb_item_batch,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品库存查询结果
	ItemInventorys []WlbItemInventory `json:"item_inventorys,omitempty" xml:"item_inventorys>wlb_item_inventory,omitempty"`
	// 商品在所有仓库的可销库存总数
	TotalQuantity int64 `json:"total_quantity,omitempty" xml:"total_quantity,omitempty"`
}
