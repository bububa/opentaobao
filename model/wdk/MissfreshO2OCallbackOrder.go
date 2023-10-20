package wdk

import (
	"sync"
)

// MissfreshO2OCallbackOrder 结构体
type MissfreshO2OCallbackOrder struct {
	// 作业单元
	CallbackUnits []MissfreshO2OCallbackUnit `json:"callback_units,omitempty" xml:"callback_units>missfresh_o2o_callback_unit,omitempty"`
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
	// 操作员
	Operator *Operator `json:"operator,omitempty" xml:"operator,omitempty"`
	// 是否作业节点终态
	IsFinal bool `json:"is_final,omitempty" xml:"is_final,omitempty"`
}

var poolMissfreshO2OCallbackOrder = sync.Pool{
	New: func() any {
		return new(MissfreshO2OCallbackOrder)
	},
}

// GetMissfreshO2OCallbackOrder() 从对象池中获取MissfreshO2OCallbackOrder
func GetMissfreshO2OCallbackOrder() *MissfreshO2OCallbackOrder {
	return poolMissfreshO2OCallbackOrder.Get().(*MissfreshO2OCallbackOrder)
}

// ReleaseMissfreshO2OCallbackOrder 释放MissfreshO2OCallbackOrder
func ReleaseMissfreshO2OCallbackOrder(v *MissfreshO2OCallbackOrder) {
	v.CallbackUnits = v.CallbackUnits[:0]
	v.Containers = v.Containers[:0]
	v.StatusChangeTime = ""
	v.StatusChangeType = ""
	v.NodeCode = ""
	v.WorkOrderType = ""
	v.WorkOrderId = ""
	v.Operator = nil
	v.IsFinal = false
	poolMissfreshO2OCallbackOrder.Put(v)
}
