package inventory

// PlanTopDto 结构体
type PlanTopDto struct {
	// 计划详情
	PlanInstanceList []PlanInstanceTopDto `json:"plan_instance_list,omitempty" xml:"plan_instance_list>plan_instance_top_dto,omitempty"`
	// 品id，前端宝贝id，或者后端货品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 品的类型，1-前端宝贝，2-后端货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 设置计划的品的skuid。对于货品，是0.
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
