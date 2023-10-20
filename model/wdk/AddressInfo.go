package wdk

import (
	"sync"
)

// AddressInfo 结构体
type AddressInfo struct {
	// 地址类型
	AddressType string `json:"address_type,omitempty" xml:"address_type,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省份
	State string `json:"state,omitempty" xml:"state,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
}

var poolAddressInfo = sync.Pool{
	New: func() any {
		return new(AddressInfo)
	},
}

// GetAddressInfo() 从对象池中获取AddressInfo
func GetAddressInfo() *AddressInfo {
	return poolAddressInfo.Get().(*AddressInfo)
}

// ReleaseAddressInfo 释放AddressInfo
func ReleaseAddressInfo(v *AddressInfo) {
	v.AddressType = ""
	v.City = ""
	v.State = ""
	v.Country = ""
	v.Address = ""
	poolAddressInfo.Put(v)
}
