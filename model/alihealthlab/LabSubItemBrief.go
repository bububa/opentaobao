package alihealthlab

import (
	"sync"
)

// LabSubItemBrief 结构体
type LabSubItemBrief struct {
	// 子项目名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 子项目isv侧编码
	IsvItemCode string `json:"isv_item_code,omitempty" xml:"isv_item_code,omitempty"`
}

var poolLabSubItemBrief = sync.Pool{
	New: func() any {
		return new(LabSubItemBrief)
	},
}

// GetLabSubItemBrief() 从对象池中获取LabSubItemBrief
func GetLabSubItemBrief() *LabSubItemBrief {
	return poolLabSubItemBrief.Get().(*LabSubItemBrief)
}

// ReleaseLabSubItemBrief 释放LabSubItemBrief
func ReleaseLabSubItemBrief(v *LabSubItemBrief) {
	v.Name = ""
	v.IsvItemCode = ""
	poolLabSubItemBrief.Put(v)
}
