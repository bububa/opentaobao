package tbitem

import (
	"sync"
)

// Location 结构体
type Location struct {
	// 所在城市（中文名称）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 所在省份（中文名称）
	State string `json:"state,omitempty" xml:"state,omitempty"`
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
	v.City = ""
	v.State = ""
	poolLocation.Put(v)
}
