package alihouse

// ProjectTradeOrderDto 结构体
type ProjectTradeOrderDto struct {
	// 楼盘列表
	ProjectList []ProjectOrderDto `json:"project_list,omitempty" xml:"project_list>project_order_dto,omitempty"`
	// 选品ID
	SelectId string `json:"select_id,omitempty" xml:"select_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 是否默认
	IsDefault int64 `json:"is_default,omitempty" xml:"is_default,omitempty"`
}
