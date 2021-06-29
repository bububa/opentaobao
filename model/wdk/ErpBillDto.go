package wdk

// ErpBillDto 
type ErpBillDto struct {
    // orderCode
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
    // status
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // createDate
    CreateDate   string `json:"create_date,omitempty" xml:"create_date,omitempty"`
    // warehouseCode
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // type
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
}
