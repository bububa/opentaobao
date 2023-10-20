package logistic

import (
	"sync"
)

// Location 结构体
type Location struct {
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
	v.Longitude = ""
	v.Latitude = ""
	poolLocation.Put(v)
}
