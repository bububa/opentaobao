package ascp

// WarehouseCooperationQueryResultDetail 结构体
type WarehouseCooperationQueryResultDetail struct {
	// 查询到的合作关系的商家仓列表
	CooperationWarehouses []CooperationWarehouse `json:"cooperation_warehouses,omitempty" xml:"cooperation_warehouses>cooperation_warehouse,omitempty"`
	// 第几页
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
