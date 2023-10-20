package ascp

import (
	"sync"
)

// RegionId 结构体
type RegionId struct {
	// 经度（高德）
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 纬度（高德）
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

var poolRegionId = sync.Pool{
	New: func() any {
		return new(RegionId)
	},
}

// GetRegionId() 从对象池中获取RegionId
func GetRegionId() *RegionId {
	return poolRegionId.Get().(*RegionId)
}

// ReleaseRegionId 释放RegionId
func ReleaseRegionId(v *RegionId) {
	v.Latitude = ""
	v.Longitude = ""
	poolRegionId.Put(v)
}
