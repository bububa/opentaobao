package simba

import (
	"sync"
)

// ItemLifeCycleViewVo 结构体
type ItemLifeCycleViewVo struct {
	// 周期文案
	LifeCycleDesc string `json:"life_cycle_desc,omitempty" xml:"life_cycle_desc,omitempty"`
	// 文案颜色
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// 周期提示文案
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
}

var poolItemLifeCycleViewVo = sync.Pool{
	New: func() any {
		return new(ItemLifeCycleViewVo)
	},
}

// GetItemLifeCycleViewVo() 从对象池中获取ItemLifeCycleViewVo
func GetItemLifeCycleViewVo() *ItemLifeCycleViewVo {
	return poolItemLifeCycleViewVo.Get().(*ItemLifeCycleViewVo)
}

// ReleaseItemLifeCycleViewVo 释放ItemLifeCycleViewVo
func ReleaseItemLifeCycleViewVo(v *ItemLifeCycleViewVo) {
	v.LifeCycleDesc = ""
	v.Color = ""
	v.Tips = ""
	poolItemLifeCycleViewVo.Put(v)
}
