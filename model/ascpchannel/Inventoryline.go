package ascpchannel

// Inventoryline 结构体
type Inventoryline struct {
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
