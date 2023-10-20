package admarket

import (
	"sync"
)

// AdSlot 结构体
type AdSlot struct {
	// 广告位id
	AdSlotId string `json:"ad_slot_id,omitempty" xml:"ad_slot_id,omitempty"`
	// 查询条件
	Query string `json:"query,omitempty" xml:"query,omitempty"`
	// 个数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolAdSlot = sync.Pool{
	New: func() any {
		return new(AdSlot)
	},
}

// GetAdSlot() 从对象池中获取AdSlot
func GetAdSlot() *AdSlot {
	return poolAdSlot.Get().(*AdSlot)
}

// ReleaseAdSlot 释放AdSlot
func ReleaseAdSlot(v *AdSlot) {
	v.AdSlotId = ""
	v.Query = ""
	v.Count = 0
	poolAdSlot.Put(v)
}
