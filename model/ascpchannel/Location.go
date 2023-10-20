package ascpchannel

import (
	"sync"
)

// Location 结构体
type Location struct {
	// 仓库编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
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
	v.StoreCode = ""
	poolLocation.Put(v)
}
