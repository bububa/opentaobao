package inventory

import (
	"sync"
)

// PlanInvTopDto 结构体
type PlanInvTopDto struct {
	// 实例列表
	PlanInstanceInvList []PlanInstanceInvResultDto `json:"plan_instance_inv_list,omitempty" xml:"plan_instance_inv_list>plan_instance_inv_result_dto,omitempty"`
}

var poolPlanInvTopDto = sync.Pool{
	New: func() any {
		return new(PlanInvTopDto)
	},
}

// GetPlanInvTopDto() 从对象池中获取PlanInvTopDto
func GetPlanInvTopDto() *PlanInvTopDto {
	return poolPlanInvTopDto.Get().(*PlanInvTopDto)
}

// ReleasePlanInvTopDto 释放PlanInvTopDto
func ReleasePlanInvTopDto(v *PlanInvTopDto) {
	v.PlanInstanceInvList = v.PlanInstanceInvList[:0]
	poolPlanInvTopDto.Put(v)
}
