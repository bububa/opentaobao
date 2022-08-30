package ascp

// ChannelInventory 结构体
type ChannelInventory struct {
	// 渠道，默认1000代表淘系渠道，预留后面可能猫超等渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 货主编码；优仓分配，长度不会超过32位，不含特殊字符
	OwnerCode string `json:"owner_code,omitempty" xml:"owner_code,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 商家货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 库存数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
