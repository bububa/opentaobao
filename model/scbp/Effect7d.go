package scbp

import (
	"sync"
)

// Effect7d 结构体
type Effect7d struct {
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 平均点击花费
	Cpc string `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 消耗金额，单位元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 点击量
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 曝光量
	Impr string `json:"impr,omitempty" xml:"impr,omitempty"`
}

var poolEffect7d = sync.Pool{
	New: func() any {
		return new(Effect7d)
	},
}

// GetEffect7d() 从对象池中获取Effect7d
func GetEffect7d() *Effect7d {
	return poolEffect7d.Get().(*Effect7d)
}

// ReleaseEffect7d 释放Effect7d
func ReleaseEffect7d(v *Effect7d) {
	v.Ctr = ""
	v.Cpc = ""
	v.Cost = ""
	v.Click = ""
	v.Impr = ""
	poolEffect7d.Put(v)
}
