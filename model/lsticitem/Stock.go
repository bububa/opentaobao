package lsticitem

// Stock 
type Stock struct {

    // 仓库编码
    
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    

    // 仓库名称
    
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    

    // 可售库存数量
    
    StockAmount   int64 `json:"stock_amount,omitempty" xml:"stock_amount,omitempty"`
    

}
