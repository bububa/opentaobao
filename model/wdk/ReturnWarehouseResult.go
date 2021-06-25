package wdk

// ReturnWarehouseResult 
type ReturnWarehouseResult struct {

    // 仓编码，由基础店仓维护，盒马全域统一
    WarehouseCode   string `json:"warehouse_code,omitempty"`

    // 入库单号
    ReturnWarehouseBillId   string `json:"return_warehouse_bill_id,omitempty"`

    // 子订单信息列表
    SkuInfoList   []ReverseSkuInfo `json:"sku_info_list,omitempty"`

}
