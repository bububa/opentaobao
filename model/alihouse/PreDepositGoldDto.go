package alihouse

// PreDepositGoldDto 结构体
type PreDepositGoldDto struct {
	// 外部楼盘id
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部销售活动id
	OuterSalesActivityId string `json:"outer_sales_activity_id,omitempty" xml:"outer_sales_activity_id,omitempty"`
	// 是否解除
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 预存金淘宝商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
