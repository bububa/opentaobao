package ascp

// WarehouseScItemRelation 结构体
type WarehouseScItemRelation struct {
	// 仓库编码
	ErpWarehouseCode string `json:"erp_warehouse_code,omitempty" xml:"erp_warehouse_code,omitempty"`
	// 仓库货品编码
	WarehouseScItemCode string `json:"warehouse_sc_item_code,omitempty" xml:"warehouse_sc_item_code,omitempty"`
}
