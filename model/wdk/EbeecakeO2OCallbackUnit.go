package wdk

import (
	"sync"
)

// EbeecakeO2OCallbackUnit 结构体
type EbeecakeO2OCallbackUnit struct {
	// 作业内容列表
	CallbackContents []EbeecakeO2OCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>ebeecake_o2o_callback_content,omitempty"`
	// 作业单元号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

var poolEbeecakeO2OCallbackUnit = sync.Pool{
	New: func() any {
		return new(EbeecakeO2OCallbackUnit)
	},
}

// GetEbeecakeO2OCallbackUnit() 从对象池中获取EbeecakeO2OCallbackUnit
func GetEbeecakeO2OCallbackUnit() *EbeecakeO2OCallbackUnit {
	return poolEbeecakeO2OCallbackUnit.Get().(*EbeecakeO2OCallbackUnit)
}

// ReleaseEbeecakeO2OCallbackUnit 释放EbeecakeO2OCallbackUnit
func ReleaseEbeecakeO2OCallbackUnit(v *EbeecakeO2OCallbackUnit) {
	v.CallbackContents = v.CallbackContents[:0]
	v.WorkOrderUnitId = ""
	poolEbeecakeO2OCallbackUnit.Put(v)
}
