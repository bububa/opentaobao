package wdk

import (
	"sync"
)

// DrfHalfDayCcCallbackOrder 结构体
type DrfHalfDayCcCallbackOrder struct {
	// 作业单元
	CallbackUnits []DrfHalfDayCcCallbackUnit `json:"callback_units,omitempty" xml:"callback_units>drf_half_day_cc_callback_unit,omitempty"`
	// 容器列表
	Containers []Container `json:"containers,omitempty" xml:"containers>container,omitempty"`
	// 作业状态变更时间
	StatusChangeTime string `json:"status_change_time,omitempty" xml:"status_change_time,omitempty"`
	// 作业状态变更类型： START_PICK(“开始拣货”)， PICK_FINISH(“拣货完成”)， START_PACKAGE(“开始打包”), PACKAGE _FINISH(“打包完成”);
	StatusChangeType string `json:"status_change_type,omitempty" xml:"status_change_type,omitempty"`
	// 节点编码
	NodeCode string `json:"node_code,omitempty" xml:"node_code,omitempty"`
	// 作业单类型： BATCH(&#34;批次&#34;),  ORDER(&#34;物流单
	WorkOrderType string `json:"work_order_type,omitempty" xml:"work_order_type,omitempty"`
	// 作业单号
	WorkOrderId string `json:"work_order_id,omitempty" xml:"work_order_id,omitempty"`
	// 作业单扩展属性
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 操作员
	Operator *Operator `json:"operator,omitempty" xml:"operator,omitempty"`
	// 是否作业节点终态
	IsFinal bool `json:"is_final,omitempty" xml:"is_final,omitempty"`
}

var poolDrfHalfDayCcCallbackOrder = sync.Pool{
	New: func() any {
		return new(DrfHalfDayCcCallbackOrder)
	},
}

// GetDrfHalfDayCcCallbackOrder() 从对象池中获取DrfHalfDayCcCallbackOrder
func GetDrfHalfDayCcCallbackOrder() *DrfHalfDayCcCallbackOrder {
	return poolDrfHalfDayCcCallbackOrder.Get().(*DrfHalfDayCcCallbackOrder)
}

// ReleaseDrfHalfDayCcCallbackOrder 释放DrfHalfDayCcCallbackOrder
func ReleaseDrfHalfDayCcCallbackOrder(v *DrfHalfDayCcCallbackOrder) {
	v.CallbackUnits = v.CallbackUnits[:0]
	v.Containers = v.Containers[:0]
	v.StatusChangeTime = ""
	v.StatusChangeType = ""
	v.NodeCode = ""
	v.WorkOrderType = ""
	v.WorkOrderId = ""
	v.Attribute = ""
	v.Operator = nil
	v.IsFinal = false
	poolDrfHalfDayCcCallbackOrder.Put(v)
}
