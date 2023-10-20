package admarket

import (
	"sync"
)

// Location 结构体
type Location struct {
	// 定位类型(WGS84/GCJ02/BD09)
	CoordinateType string `json:"coordinate_type,omitempty" xml:"coordinate_type,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
}

var poolLocation = sync.Pool{
	New: func() any {
		return new(Location)
	},
}

// GetLocation() 从对象池中获取Location
func GetLocation() *Location {
	return poolLocation.Get().(*Location)
}

// ReleaseLocation 释放Location
func ReleaseLocation(v *Location) {
	v.CoordinateType = ""
	v.Longitude = ""
	v.Latitude = ""
	poolLocation.Put(v)
}
