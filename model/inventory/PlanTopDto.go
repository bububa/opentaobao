package inventory

import (
	"sync"
)

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

var poolPlanTopDto = sync.Pool{
	New: func() any {
		return new(PlanTopDto)
	},
}

// GetPlanTopDto() 从对象池中获取PlanTopDto
func GetPlanTopDto() *PlanTopDto {
	return poolPlanTopDto.Get().(*PlanTopDto)
}

// ReleasePlanTopDto 释放PlanTopDto
func ReleasePlanTopDto(v *PlanTopDto) {
	v.PlanInstanceList = v.PlanInstanceList[:0]
	v.ItemId = 0
	v.ItemType = 0
	v.SkuId = 0
	poolPlanTopDto.Put(v)
}
