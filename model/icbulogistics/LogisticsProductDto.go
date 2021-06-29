package icbulogistics

// LogisticsProductDto 
type LogisticsProductDto struct {
    // 仓库名称
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    // 仓库编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 产品名称
    ProductName   string `json:"product_name,omitempty" xml:"product_name,omitempty"`
    // 产品编码
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    // 是否上门揽收
    Pickup   bool `json:"pickup,omitempty" xml:"pickup,omitempty"`
}
