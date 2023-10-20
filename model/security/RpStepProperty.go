package security

import (
	"sync"
)

// RpStepProperty 结构体
type RpStepProperty struct {
	// property
	Property *RpProperty `json:"property,omitempty" xml:"property,omitempty"`
	// optional
	Optional bool `json:"optional,omitempty" xml:"optional,omitempty"`
	// 是否可选
	IsOptional bool `json:"is_optional,omitempty" xml:"is_optional,omitempty"`
}

var poolRpStepProperty = sync.Pool{
	New: func() any {
		return new(RpStepProperty)
	},
}

// GetRpStepProperty() 从对象池中获取RpStepProperty
func GetRpStepProperty() *RpStepProperty {
	return poolRpStepProperty.Get().(*RpStepProperty)
}

// ReleaseRpStepProperty 释放RpStepProperty
func ReleaseRpStepProperty(v *RpStepProperty) {
	v.Property = nil
	v.Optional = false
	v.IsOptional = false
	poolRpStepProperty.Put(v)
}
