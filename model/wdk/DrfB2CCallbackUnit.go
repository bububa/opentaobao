package wdk

import (
	"sync"
)

// DrfB2CCallbackUnit 结构体
type DrfB2CCallbackUnit struct {
	// 作业内容
	CallbackContents []DrfB2CCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>drf_b2c_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

var poolDrfB2CCallbackUnit = sync.Pool{
	New: func() any {
		return new(DrfB2CCallbackUnit)
	},
}

// GetDrfB2CCallbackUnit() 从对象池中获取DrfB2CCallbackUnit
func GetDrfB2CCallbackUnit() *DrfB2CCallbackUnit {
	return poolDrfB2CCallbackUnit.Get().(*DrfB2CCallbackUnit)
}

// ReleaseDrfB2CCallbackUnit 释放DrfB2CCallbackUnit
func ReleaseDrfB2CCallbackUnit(v *DrfB2CCallbackUnit) {
	v.CallbackContents = v.CallbackContents[:0]
	v.WorkOrderUnitId = ""
	poolDrfB2CCallbackUnit.Put(v)
}
