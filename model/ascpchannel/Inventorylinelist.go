package ascpchannel

// Inventorylinelist 结构体
type Inventorylinelist struct {
	// 在库库存操作行对象
	InventoryLine *Inventoryline `json:"inventory_line,omitempty" xml:"inventory_line,omitempty"`
}
