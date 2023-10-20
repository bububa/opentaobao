package wdk

import (
	"sync"
)

// SfB2CFmsCallbackUnit 结构体
type SfB2CFmsCallbackUnit struct {
	// 作业内容
	CallbackContents []SfB2CFmsCallbackContent `json:"callback_contents,omitempty" xml:"callback_contents>sf_b2c_fms_callback_content,omitempty"`
	// 作业单元单号
	WorkOrderUnitId string `json:"work_order_unit_id,omitempty" xml:"work_order_unit_id,omitempty"`
}

var poolSfB2CFmsCallbackUnit = sync.Pool{
	New: func() any {
		return new(SfB2CFmsCallbackUnit)
	},
}

// GetSfB2CFmsCallbackUnit() 从对象池中获取SfB2CFmsCallbackUnit
func GetSfB2CFmsCallbackUnit() *SfB2CFmsCallbackUnit {
	return poolSfB2CFmsCallbackUnit.Get().(*SfB2CFmsCallbackUnit)
}

// ReleaseSfB2CFmsCallbackUnit 释放SfB2CFmsCallbackUnit
func ReleaseSfB2CFmsCallbackUnit(v *SfB2CFmsCallbackUnit) {
	v.CallbackContents = v.CallbackContents[:0]
	v.WorkOrderUnitId = ""
	poolSfB2CFmsCallbackUnit.Put(v)
}
