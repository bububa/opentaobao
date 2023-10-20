package tblogistics

import (
	"sync"
)

// Location 结构体
type Location struct {
	// 邮政编码
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// 详细地址，最大256个字节（128个中文）
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 所在城市（中文名称）
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 所在省份（中文名称）
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 国家名称
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 区/县（只适用于物流API）
	District string `json:"district,omitempty" xml:"district,omitempty"`
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
	v.Zip = ""
	v.Address = ""
	v.City = ""
	v.State = ""
	v.Country = ""
	v.District = ""
	poolLocation.Put(v)
}
