package kbalgo

import (
	"sync"
)

// Label 结构体
type Label struct {
	// 是否外卖
	LabelDescription string `json:"label_description,omitempty" xml:"label_description,omitempty"`
	// 标签类型
	LabelType string `json:"label_type,omitempty" xml:"label_type,omitempty"`
	// Delivery
	Delivery *Delivery `json:"delivery,omitempty" xml:"delivery,omitempty"`
	// 链接
	Schema *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
}

var poolLabel = sync.Pool{
	New: func() any {
		return new(Label)
	},
}

// GetLabel() 从对象池中获取Label
func GetLabel() *Label {
	return poolLabel.Get().(*Label)
}

// ReleaseLabel 释放Label
func ReleaseLabel(v *Label) {
	v.LabelDescription = ""
	v.LabelType = ""
	v.Delivery = nil
	v.Schema = nil
	poolLabel.Put(v)
}
