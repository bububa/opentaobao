package drugtrace

import (
	"sync"
)

// Address 结构体
type Address struct {
	// 境内填写区县名称/境外则填写境外国家中文名称
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 城市名称/境外不用填，境内必填
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 省份名称/境外不用填，境内必填
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
}

var poolAddress = sync.Pool{
	New: func() any {
		return new(Address)
	},
}

// GetAddress() 从对象池中获取Address
func GetAddress() *Address {
	return poolAddress.Get().(*Address)
}

// ReleaseAddress 释放Address
func ReleaseAddress(v *Address) {
	v.AreaName = ""
	v.CityName = ""
	v.ProvName = ""
	poolAddress.Put(v)
}
