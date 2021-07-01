package waybill

// WarehouseDto 结构体
type WarehouseDto struct {
	// 仓id
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// 仓名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
}
