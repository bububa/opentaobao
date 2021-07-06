package rhino

// MaterialInventoryAdjustDto 结构体
type MaterialInventoryAdjustDto struct {
	// 物料名称
	MaterialName string `json:"material_name,omitempty" xml:"material_name,omitempty"`
	// 库存状态(1:正品;2:次品;3:赠品)
	InventoryType string `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 调整标示(1:增加库存;2:减少库存)
	AdjustType string `json:"adjust_type,omitempty" xml:"adjust_type,omitempty"`
	// 讯息ID
	AdjustSerialNo string `json:"adjust_serial_no,omitempty" xml:"adjust_serial_no,omitempty"`
	// 盘点原因
	AdjustReason string `json:"adjust_reason,omitempty" xml:"adjust_reason,omitempty"`
	// 盘点日期
	AdjustDate string `json:"adjust_date,omitempty" xml:"adjust_date,omitempty"`
	// 工厂ID，货主编码
	FactoryId int64 `json:"factory_id,omitempty" xml:"factory_id,omitempty"`
	// 物料ID，货品编码
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 仓库ID，仓库编码
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// 盘点数量
	AdjustAmount int64 `json:"adjust_amount,omitempty" xml:"adjust_amount,omitempty"`
}
