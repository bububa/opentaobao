package ascpchannel

// Transferextendorderitemdtolist 结构体
type Transferextendorderitemdtolist struct {
	// 批次号
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 库存类型
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
