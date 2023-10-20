package wdk

import (
	"sync"
)

// WorkOrder 结构体
type WorkOrder struct {
	// 作业单列表
	WorkOrderUnits []WorkOrderUnits `json:"work_order_units,omitempty" xml:"work_order_units>work_order_units,omitempty"`
	// 仓编码
	WareHouseCode string `json:"ware_house_code,omitempty" xml:"ware_house_code,omitempty"`
	// 作业单最早送达时间
	EarliestArrivalTime string `json:"earliest_arrival_time,omitempty" xml:"earliest_arrival_time,omitempty"`
	// 任务名称
	WorkOrderName string `json:"work_order_name,omitempty" xml:"work_order_name,omitempty"`
	// 作业单最晚送达时间
	LatestArriveTime string `json:"latest_arrive_time,omitempty" xml:"latest_arrive_time,omitempty"`
	// 任务编码
	WorkOrderId string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
}

var poolWorkOrder = sync.Pool{
	New: func() any {
		return new(WorkOrder)
	},
}

// GetWorkOrder() 从对象池中获取WorkOrder
func GetWorkOrder() *WorkOrder {
	return poolWorkOrder.Get().(*WorkOrder)
}

// ReleaseWorkOrder 释放WorkOrder
func ReleaseWorkOrder(v *WorkOrder) {
	v.WorkOrderUnits = v.WorkOrderUnits[:0]
	v.WareHouseCode = ""
	v.EarliestArrivalTime = ""
	v.WorkOrderName = ""
	v.LatestArriveTime = ""
	v.WorkOrderId = ""
	poolWorkOrder.Put(v)
}
