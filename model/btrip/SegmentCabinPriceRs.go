package btrip

import (
	"sync"
)

// SegmentCabinPriceRs 结构体
type SegmentCabinPriceRs struct {
	// 仓位信息
	Cabin *CabinRs `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 搜索提供的价格信息
	SearchPrice *SearchPriceRs `json:"search_price,omitempty" xml:"search_price,omitempty"`
	// 航段
	Segment *SegmentPositionRs `json:"segment,omitempty" xml:"segment,omitempty"`
}

var poolSegmentCabinPriceRs = sync.Pool{
	New: func() any {
		return new(SegmentCabinPriceRs)
	},
}

// GetSegmentCabinPriceRs() 从对象池中获取SegmentCabinPriceRs
func GetSegmentCabinPriceRs() *SegmentCabinPriceRs {
	return poolSegmentCabinPriceRs.Get().(*SegmentCabinPriceRs)
}

// ReleaseSegmentCabinPriceRs 释放SegmentCabinPriceRs
func ReleaseSegmentCabinPriceRs(v *SegmentCabinPriceRs) {
	v.Cabin = nil
	v.SearchPrice = nil
	v.Segment = nil
	poolSegmentCabinPriceRs.Put(v)
}
