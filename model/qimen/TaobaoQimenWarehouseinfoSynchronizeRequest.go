package qimen

// TaobaoqimenwarehouseinfosynchronizeRequest 结构体
type TaobaoqimenwarehouseinfosynchronizeRequest struct {
	// 仓库信息
	WarehouseInfos []WarehouseInfos `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_infos,omitempty"`
	// 货主编码，string（50）
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 货主名称，string（50）
	OwnerName string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
}
