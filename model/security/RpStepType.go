package security

import (
	"sync"
)

// RpStepType 结构体
type RpStepType struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpStepType = sync.Pool{
	New: func() any {
		return new(RpStepType)
	},
}

// GetRpStepType() 从对象池中获取RpStepType
func GetRpStepType() *RpStepType {
	return poolRpStepType.Get().(*RpStepType)
}

// ReleaseRpStepType 释放RpStepType
func ReleaseRpStepType(v *RpStepType) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpStepType.Put(v)
}
