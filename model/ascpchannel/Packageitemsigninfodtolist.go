package ascpchannel

// Packageitemsigninfodtolist 结构体
type Packageitemsigninfodtolist struct {
	// 签收状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}
