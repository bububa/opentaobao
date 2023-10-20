package wdk

import (
	"sync"
)

// DrfHalfDayCcCallbackUnit 结构体
type DrfHalfDayCcCallbackUnit struct {
	// 作业内容
	CallbackContents []DrfHalfDayCcCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>drf_half_day_cc_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
	// 作业单元扩展属性
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
}

var poolDrfHalfDayCcCallbackUnit = sync.Pool{
	New: func() any {
		return new(DrfHalfDayCcCallbackUnit)
	},
}

// GetDrfHalfDayCcCallbackUnit() 从对象池中获取DrfHalfDayCcCallbackUnit
func GetDrfHalfDayCcCallbackUnit() *DrfHalfDayCcCallbackUnit {
	return poolDrfHalfDayCcCallbackUnit.Get().(*DrfHalfDayCcCallbackUnit)
}

// ReleaseDrfHalfDayCcCallbackUnit 释放DrfHalfDayCcCallbackUnit
func ReleaseDrfHalfDayCcCallbackUnit(v *DrfHalfDayCcCallbackUnit) {
	v.CallbackContents = v.CallbackContents[:0]
	v.WorkOrderUnitId = ""
	v.Attribute = ""
	poolDrfHalfDayCcCallbackUnit.Put(v)
}
