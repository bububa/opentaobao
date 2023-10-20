package btrip

import (
	"sync"
)

// TripPreferenceRq 结构体
type TripPreferenceRq struct {
	// 仓位代码
	CabinList []string `json:"cabin_list,omitempty" xml:"cabin_list>string,omitempty"`
}

var poolTripPreferenceRq = sync.Pool{
	New: func() any {
		return new(TripPreferenceRq)
	},
}

// GetTripPreferenceRq() 从对象池中获取TripPreferenceRq
func GetTripPreferenceRq() *TripPreferenceRq {
	return poolTripPreferenceRq.Get().(*TripPreferenceRq)
}

// ReleaseTripPreferenceRq 释放TripPreferenceRq
func ReleaseTripPreferenceRq(v *TripPreferenceRq) {
	v.CabinList = v.CabinList[:0]
	poolTripPreferenceRq.Put(v)
}
