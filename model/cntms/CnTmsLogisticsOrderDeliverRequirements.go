package cntms

import (
	"sync"
)

// CnTmsLogisticsOrderDeliverRequirements 结构体
type CnTmsLogisticsOrderDeliverRequirements struct {
	// 配送类型： PTPS-普通配送 LLPS-冷链配送
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// 投递时延要求(1 工作日 2 节假日 104 预约达 )
	ScheduleType string `json:"schedule_type,omitempty" xml:"schedule_type,omitempty"`
	// 送达日期（格式为 yyyy-MM-dd)
	ScheduleDay string `json:"schedule_day,omitempty" xml:"schedule_day,omitempty"`
	// 送达开始时间（格式为 hh:mm）
	ScheduleStart string `json:"schedule_start,omitempty" xml:"schedule_start,omitempty"`
	// 送达结束时间（格式为 hh:mm）
	ScheduleEnd string `json:"schedule_end,omitempty" xml:"schedule_end,omitempty"`
}

var poolCnTmsLogisticsOrderDeliverRequirements = sync.Pool{
	New: func() any {
		return new(CnTmsLogisticsOrderDeliverRequirements)
	},
}

// GetCnTmsLogisticsOrderDeliverRequirements() 从对象池中获取CnTmsLogisticsOrderDeliverRequirements
func GetCnTmsLogisticsOrderDeliverRequirements() *CnTmsLogisticsOrderDeliverRequirements {
	return poolCnTmsLogisticsOrderDeliverRequirements.Get().(*CnTmsLogisticsOrderDeliverRequirements)
}

// ReleaseCnTmsLogisticsOrderDeliverRequirements 释放CnTmsLogisticsOrderDeliverRequirements
func ReleaseCnTmsLogisticsOrderDeliverRequirements(v *CnTmsLogisticsOrderDeliverRequirements) {
	v.DeliveryType = ""
	v.ScheduleType = ""
	v.ScheduleDay = ""
	v.ScheduleStart = ""
	v.ScheduleEnd = ""
	poolCnTmsLogisticsOrderDeliverRequirements.Put(v)
}
