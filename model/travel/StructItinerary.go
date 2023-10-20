package travel

import (
	"sync"
)

// StructItinerary 结构体
type StructItinerary struct {
	// 必填，当天行程包含的多个活动信息
	Activities []ItineraryActivity `json:"activities,omitempty" xml:"activities>itinerary_activity,omitempty"`
	// 必填，行程序号，标识是第几天的行程
	DayOrder int64 `json:"day_order,omitempty" xml:"day_order,omitempty"`
}

var poolStructItinerary = sync.Pool{
	New: func() any {
		return new(StructItinerary)
	},
}

// GetStructItinerary() 从对象池中获取StructItinerary
func GetStructItinerary() *StructItinerary {
	return poolStructItinerary.Get().(*StructItinerary)
}

// ReleaseStructItinerary 释放StructItinerary
func ReleaseStructItinerary(v *StructItinerary) {
	v.Activities = v.Activities[:0]
	v.DayOrder = 0
	poolStructItinerary.Put(v)
}
