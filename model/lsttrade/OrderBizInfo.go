package lsttrade

// OrderBizInfo 
type OrderBizInfo struct {
    // 仓库名称
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 零售通仓库类型。customer：虚仓；cainiao：实仓
    LstWarehouseType   string `json:"lst_warehouse_type,omitempty" xml:"lst_warehouse_type,omitempty"`
}
