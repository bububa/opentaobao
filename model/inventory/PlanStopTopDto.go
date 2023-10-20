package inventory

import (
	"sync"
)

// PlanStopTopDto 结构体
type PlanStopTopDto struct {
	// 要失效的计划库存的详情
	PlanDetailList []PlanStopDetailTopDto `json:"plan_detail_list,omitempty" xml:"plan_detail_list>plan_stop_detail_top_dto,omitempty"`
	// 品id，前端宝贝id，或者后端货品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 品的类型，1-前端宝贝，2-后端货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 设置计划的品的skuid。对于货品，是0.
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolPlanStopTopDto = sync.Pool{
	New: func() any {
		return new(PlanStopTopDto)
	},
}

// GetPlanStopTopDto() 从对象池中获取PlanStopTopDto
func GetPlanStopTopDto() *PlanStopTopDto {
	return poolPlanStopTopDto.Get().(*PlanStopTopDto)
}

// ReleasePlanStopTopDto 释放PlanStopTopDto
func ReleasePlanStopTopDto(v *PlanStopTopDto) {
	v.PlanDetailList = v.PlanDetailList[:0]
	v.ItemId = 0
	v.ItemType = 0
	v.SkuId = 0
	poolPlanStopTopDto.Put(v)
}
