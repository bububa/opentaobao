package qimen

// ItemsSynchronizeRequest 结构体
type ItemsSynchronizeRequest struct {
	// 同步的商品信息
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 操作类型(两种类型：add|update)
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenItemsSynchronizeMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
