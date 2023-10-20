package logistic

import (
	"sync"
)

// ReceiveAddress 结构体
type ReceiveAddress struct {
	// 镇/街道
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区
	AreaName string `json:"area_name,omitempty" xml:"area_name,omitempty"`
	// 省
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
}

var poolReceiveAddress = sync.Pool{
	New: func() any {
		return new(ReceiveAddress)
	},
}

// GetReceiveAddress() 从对象池中获取ReceiveAddress
func GetReceiveAddress() *ReceiveAddress {
	return poolReceiveAddress.Get().(*ReceiveAddress)
}

// ReleaseReceiveAddress 释放ReceiveAddress
func ReleaseReceiveAddress(v *ReceiveAddress) {
	v.TownName = ""
	v.AddressDetail = ""
	v.CityName = ""
	v.AreaName = ""
	v.ProvinceName = ""
	poolReceiveAddress.Put(v)
}
