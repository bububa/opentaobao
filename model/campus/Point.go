package campus

import (
	"sync"
)

// Point 结构体
type Point struct {
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
}

var poolPoint = sync.Pool{
	New: func() any {
		return new(Point)
	},
}

// GetPoint() 从对象池中获取Point
func GetPoint() *Point {
	return poolPoint.Get().(*Point)
}

// ReleasePoint 释放Point
func ReleasePoint(v *Point) {
	v.Lng = ""
	v.Lat = ""
	poolPoint.Put(v)
}
