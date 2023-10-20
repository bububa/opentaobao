package inventory

import (
	"sync"
)

// PlanPromiseTopDto 结构体
type PlanPromiseTopDto struct {
	// 仓code
	PerformStore string `json:"perform_store,omitempty" xml:"perform_store,omitempty"`
	// 履约时间，如绝对时间 2021-03-11 ，或者相对时间 2
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
	// 履约时间类型，0 是绝对时间，年月日；1是相对时间
	DeliveryType int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
}

var poolPlanPromiseTopDto = sync.Pool{
	New: func() any {
		return new(PlanPromiseTopDto)
	},
}

// GetPlanPromiseTopDto() 从对象池中获取PlanPromiseTopDto
func GetPlanPromiseTopDto() *PlanPromiseTopDto {
	return poolPlanPromiseTopDto.Get().(*PlanPromiseTopDto)
}

// ReleasePlanPromiseTopDto 释放PlanPromiseTopDto
func ReleasePlanPromiseTopDto(v *PlanPromiseTopDto) {
	v.PerformStore = ""
	v.DeliveryTime = ""
	v.DeliveryType = 0
	poolPlanPromiseTopDto.Put(v)
}
