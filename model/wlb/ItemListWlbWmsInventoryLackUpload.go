package wlb

// ItemListWlbWmsInventoryLackUpload 结构体
type ItemListWlbWmsInventoryLackUpload struct {
	// 报缺原因（系统报缺，实物报缺）
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 后端商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 应发商品数量
	ItemQty int64 `json:"item_qty,omitempty" xml:"item_qty,omitempty"`
	// 商品缺货数量
	LackQty int64 `json:"lack_qty,omitempty" xml:"lack_qty,omitempty"`
}
