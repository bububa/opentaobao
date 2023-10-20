package icbulogistics

import (
	"sync"
)

// Town 结构体
type Town struct {
	// 地址代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地址名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolTown = sync.Pool{
	New: func() any {
		return new(Town)
	},
}

// GetTown() 从对象池中获取Town
func GetTown() *Town {
	return poolTown.Get().(*Town)
}

// ReleaseTown 释放Town
func ReleaseTown(v *Town) {
	v.Code = ""
	v.Name = ""
	v.AreaId = 0
	poolTown.Put(v)
}
