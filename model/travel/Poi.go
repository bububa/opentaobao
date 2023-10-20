package travel

import (
	"sync"
)

// Poi 结构体
type Poi struct {
	// POI对应的名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// POI对应ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolPoi = sync.Pool{
	New: func() any {
		return new(Poi)
	},
}

// GetPoi() 从对象池中获取Poi
func GetPoi() *Poi {
	return poolPoi.Get().(*Poi)
}

// ReleasePoi 释放Poi
func ReleasePoi(v *Poi) {
	v.Name = ""
	v.Id = 0
	poolPoi.Put(v)
}
