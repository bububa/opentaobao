package inventory

// InventoryCheckDto 结构体
type InventoryCheckDto struct {
	// 1:全量更新   2: 出入库盘盈盘亏
	CheckMode int64 `json:"check_mode,omitempty" xml:"check_mode,omitempty"`
	// 2： 仓库类型   6：门店类型
	InvStoreType int64 `json:"inv_store_type,omitempty" xml:"inv_store_type,omitempty"`
	// 调整明细
	DetailList []InventoryCheckDetailDto `json:"detail_list,omitempty" xml:"detail_list>inventory_check_detail_dto,omitempty"`
	// 仓库code或者门店id
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 调整单据号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
