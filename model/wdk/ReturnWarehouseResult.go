package wdk

// ReturnWarehouseResult 结构体
type ReturnWarehouseResult struct {
	// 子订单信息列表
	SkuInfoList []ReverseSkuInfo `json:"sku_info_list,omitempty" xml:"sku_info_list>reverse_sku_info,omitempty"`
	// 仓编码，由基础店仓维护，盒马全域统一
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 入库单号
	ReturnWarehouseBillId string `json:"return_warehouse_bill_id,omitempty" xml:"return_warehouse_bill_id,omitempty"`
	// 0:包裹完整 1:包裹破损
	PackageQuality string `json:"package_quality,omitempty" xml:"package_quality,omitempty"`
}
