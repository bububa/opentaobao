package wdk

// ErpReturnItemDto 
type ErpReturnItemDto struct {
    // 数量
    Count   string `json:"count,omitempty" xml:"count,omitempty"`
    // 库位号，退货库位号
    CabinetCode   string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
    // 商品code
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    // warehouseCode
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}
