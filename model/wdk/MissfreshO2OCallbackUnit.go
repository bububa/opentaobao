package wdk

import (
	"sync"
)

// MissfreshO2OCallbackUnit 结构体
type MissfreshO2OCallbackUnit struct {
	// 作业内容
	CallbackContents []MissfreshO2OCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>missfresh_o2o_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

var poolMissfreshO2OCallbackUnit = sync.Pool{
	New: func() any {
		return new(MissfreshO2OCallbackUnit)
	},
}

// GetMissfreshO2OCallbackUnit() 从对象池中获取MissfreshO2OCallbackUnit
func GetMissfreshO2OCallbackUnit() *MissfreshO2OCallbackUnit {
	return poolMissfreshO2OCallbackUnit.Get().(*MissfreshO2OCallbackUnit)
}

// ReleaseMissfreshO2OCallbackUnit 释放MissfreshO2OCallbackUnit
func ReleaseMissfreshO2OCallbackUnit(v *MissfreshO2OCallbackUnit) {
	v.CallbackContents = v.CallbackContents[:0]
	v.WorkOrderUnitId = ""
	poolMissfreshO2OCallbackUnit.Put(v)
}
