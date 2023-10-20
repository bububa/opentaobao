package admarket

import (
	"sync"
)

// AdSlots 结构体
type AdSlots struct {
	// 广告集合
	Ads []Ad `json:"ads,omitempty" xml:"ads>ad,omitempty"`
	// 广告位id
	AdSlotId string `json:"ad_slot_id,omitempty" xml:"ad_slot_id,omitempty"`
}

var poolAdSlots = sync.Pool{
	New: func() any {
		return new(AdSlots)
	},
}

// GetAdSlots() 从对象池中获取AdSlots
func GetAdSlots() *AdSlots {
	return poolAdSlots.Get().(*AdSlots)
}

// ReleaseAdSlots 释放AdSlots
func ReleaseAdSlots(v *AdSlots) {
	v.Ads = v.Ads[:0]
	v.AdSlotId = ""
	poolAdSlots.Put(v)
}
