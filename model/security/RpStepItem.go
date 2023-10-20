package security

import (
	"sync"
)

// RpStepItem 结构体
type RpStepItem struct {
	// properties
	Properties []RpStepProperty `json:"properties,omitempty" xml:"properties>rp_step_property,omitempty"`
	// jsonAssist
	JsonAssist string `json:"json_assist,omitempty" xml:"json_assist,omitempty"`
	// stepType
	StepType *RpStepType `json:"step_type,omitempty" xml:"step_type,omitempty"`
}

var poolRpStepItem = sync.Pool{
	New: func() any {
		return new(RpStepItem)
	},
}

// GetRpStepItem() 从对象池中获取RpStepItem
func GetRpStepItem() *RpStepItem {
	return poolRpStepItem.Get().(*RpStepItem)
}

// ReleaseRpStepItem 释放RpStepItem
func ReleaseRpStepItem(v *RpStepItem) {
	v.Properties = v.Properties[:0]
	v.JsonAssist = ""
	v.StepType = nil
	poolRpStepItem.Put(v)
}
