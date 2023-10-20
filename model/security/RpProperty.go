package security

import (
	"sync"
)

// RpProperty 结构体
type RpProperty struct {
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpProperty = sync.Pool{
	New: func() any {
		return new(RpProperty)
	},
}

// GetRpProperty() 从对象池中获取RpProperty
func GetRpProperty() *RpProperty {
	return poolRpProperty.Get().(*RpProperty)
}

// ReleaseRpProperty 释放RpProperty
func ReleaseRpProperty(v *RpProperty) {
	v.Name = ""
	v.Code = 0
	poolRpProperty.Put(v)
}
