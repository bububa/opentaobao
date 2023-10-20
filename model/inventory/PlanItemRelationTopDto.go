package inventory

import (
	"sync"
)

// PlanItemRelationTopDto 结构体
type PlanItemRelationTopDto struct {
	// 计划生效的itemid
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 计划生效的skuid
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

var poolPlanItemRelationTopDto = sync.Pool{
	New: func() any {
		return new(PlanItemRelationTopDto)
	},
}

// GetPlanItemRelationTopDto() 从对象池中获取PlanItemRelationTopDto
func GetPlanItemRelationTopDto() *PlanItemRelationTopDto {
	return poolPlanItemRelationTopDto.Get().(*PlanItemRelationTopDto)
}

// ReleasePlanItemRelationTopDto 释放PlanItemRelationTopDto
func ReleasePlanItemRelationTopDto(v *PlanItemRelationTopDto) {
	v.ItemId = 0
	v.SkuId = 0
	poolPlanItemRelationTopDto.Put(v)
}
