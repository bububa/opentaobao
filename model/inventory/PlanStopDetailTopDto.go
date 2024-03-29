package inventory

import (
	"sync"
)

// PlanStopDetailTopDto 结构体
type PlanStopDetailTopDto struct {
	// 外部商家系统单据号，用于生成计划库存的业务来源
	PlanOrderId string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
	// 操作码
	OperateCode string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
}

var poolPlanStopDetailTopDto = sync.Pool{
	New: func() any {
		return new(PlanStopDetailTopDto)
	},
}

// GetPlanStopDetailTopDto() 从对象池中获取PlanStopDetailTopDto
func GetPlanStopDetailTopDto() *PlanStopDetailTopDto {
	return poolPlanStopDetailTopDto.Get().(*PlanStopDetailTopDto)
}

// ReleasePlanStopDetailTopDto 释放PlanStopDetailTopDto
func ReleasePlanStopDetailTopDto(v *PlanStopDetailTopDto) {
	v.PlanOrderId = ""
	v.OperateCode = ""
	poolPlanStopDetailTopDto.Put(v)
}
