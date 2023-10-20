package ascp

import (
	"sync"
)

// RegionIds 结构体
type RegionIds struct {
	// 经度（高德）
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度（高德）
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
}

var poolRegionIds = sync.Pool{
	New: func() any {
		return new(RegionIds)
	},
}

// GetRegionIds() 从对象池中获取RegionIds
func GetRegionIds() *RegionIds {
	return poolRegionIds.Get().(*RegionIds)
}

// ReleaseRegionIds 释放RegionIds
func ReleaseRegionIds(v *RegionIds) {
	v.Longitude = ""
	v.Latitude = ""
	poolRegionIds.Put(v)
}
