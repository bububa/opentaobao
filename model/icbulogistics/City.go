package icbulogistics

import (
	"sync"
)

// City 结构体
type City struct {
	// 地址代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地址名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址id
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
}

var poolCity = sync.Pool{
	New: func() any {
		return new(City)
	},
}

// GetCity() 从对象池中获取City
func GetCity() *City {
	return poolCity.Get().(*City)
}

// ReleaseCity 释放City
func ReleaseCity(v *City) {
	v.Code = ""
	v.Name = ""
	v.AreaId = 0
	poolCity.Put(v)
}
