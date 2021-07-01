package inventory

// PlanStopTopDto 结构体
type PlanStopTopDto struct {
	// 品id，前端宝贝id，或者后端货品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 品的类型，1-前端宝贝，2-后端货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 要失效的计划库存的详情
	PlanDetailList []PlanStopDetailTopDto `json:"plan_detail_list,omitempty" xml:"plan_detail_list>plan_stop_detail_top_dto,omitempty"`
	// 设置计划的品的skuid。对于货品，是0.
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
