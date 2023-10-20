package inventory

import (
	"sync"
)

// PlanInstanceTopDto 结构体
type PlanInstanceTopDto struct {
	// 服务承诺信息
	PromiseList []PromiseList `json:"promise_list,omitempty" xml:"promise_list>promise_list,omitempty"`
	// 计划库存生效的宝贝范围。对于品的后端货品情况下，才有可能设置，设置后在这个范围的宝贝才能使用计划库存。也可以不设置，代表后端货品关联的所有宝贝都生效。
	RelationList []RelationList `json:"relation_list,omitempty" xml:"relation_list>relation_list,omitempty"`
	// 计划库存销售开始时间，年月日时分秒
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 计划库存销售结束时间，年月日时分秒
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 外部商家系统单据号，用于生成计划库存的业务来源
	PlanOrderId string `json:"plan_order_id,omitempty" xml:"plan_order_id,omitempty"`
	// 操作码，用于幂等验证
	OperateCode string `json:"operate_code,omitempty" xml:"operate_code,omitempty"`
	// 设置的库存类型，0-全量覆盖，1-增量处理
	QuantityOpType int64 `json:"quantity_op_type,omitempty" xml:"quantity_op_type,omitempty"`
	// 销售策略
	Strategy *Strategy `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// 设置的库存值
	SettingQuantity int64 `json:"setting_quantity,omitempty" xml:"setting_quantity,omitempty"`
}

var poolPlanInstanceTopDto = sync.Pool{
	New: func() any {
		return new(PlanInstanceTopDto)
	},
}

// GetPlanInstanceTopDto() 从对象池中获取PlanInstanceTopDto
func GetPlanInstanceTopDto() *PlanInstanceTopDto {
	return poolPlanInstanceTopDto.Get().(*PlanInstanceTopDto)
}

// ReleasePlanInstanceTopDto 释放PlanInstanceTopDto
func ReleasePlanInstanceTopDto(v *PlanInstanceTopDto) {
	v.PromiseList = v.PromiseList[:0]
	v.RelationList = v.RelationList[:0]
	v.StartTime = ""
	v.EndTime = ""
	v.PlanOrderId = ""
	v.OperateCode = ""
	v.QuantityOpType = 0
	v.Strategy = nil
	v.SettingQuantity = 0
	poolPlanInstanceTopDto.Put(v)
}
