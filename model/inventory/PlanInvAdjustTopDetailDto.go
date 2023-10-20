package inventory

import (
	"sync"
)

// PlanInvAdjustTopDetailDto 结构体
type PlanInvAdjustTopDetailDto struct {
	// 操作码，用于幂等验证
	OperateCode string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
	// 外部商家系统单据号，用于定位计划库存的业务来源
	PlanOrderId string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
	// 要调整的库存值，负数代表调低库存，正数代表调高库存
	AdjustQuantity int64 `json:"adjust_quantity,omitempty" xml:"adjust_quantity,omitempty"`
}

var poolPlanInvAdjustTopDetailDto = sync.Pool{
	New: func() any {
		return new(PlanInvAdjustTopDetailDto)
	},
}

// GetPlanInvAdjustTopDetailDto() 从对象池中获取PlanInvAdjustTopDetailDto
func GetPlanInvAdjustTopDetailDto() *PlanInvAdjustTopDetailDto {
	return poolPlanInvAdjustTopDetailDto.Get().(*PlanInvAdjustTopDetailDto)
}

// ReleasePlanInvAdjustTopDetailDto 释放PlanInvAdjustTopDetailDto
func ReleasePlanInvAdjustTopDetailDto(v *PlanInvAdjustTopDetailDto) {
	v.OperateCode = ""
	v.PlanOrderId = ""
	v.AdjustQuantity = 0
	poolPlanInvAdjustTopDetailDto.Put(v)
}
