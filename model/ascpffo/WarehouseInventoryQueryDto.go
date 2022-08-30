package ascpffo

// WarehouseInventoryQueryDto 结构体
type WarehouseInventoryQueryDto struct {
	// 货品列表，最多30个
	ScItemIdList []int64 `json:"sc_item_id_list,omitempty" xml:"sc_item_id_list>int64,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 库存类型（1 良品，101 残品）
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小，最大30
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
