package tmallnr

import (
	"sync"
)

// Points 结构体
type Points struct {
	// 维度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
}

var poolPoints = sync.Pool{
	New: func() any {
		return new(Points)
	},
}

// GetPoints() 从对象池中获取Points
func GetPoints() *Points {
	return poolPoints.Get().(*Points)
}

// ReleasePoints 释放Points
func ReleasePoints(v *Points) {
	v.Lat = ""
	v.Lng = ""
	poolPoints.Put(v)
}
