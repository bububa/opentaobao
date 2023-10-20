package wlb

import (
	"sync"
)

// AddressInfo 结构体
type AddressInfo struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Borough string `json:"borough,omitempty" xml:"borough,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 邮编
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
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
	v.Province = ""
	v.City = ""
	v.Borough = ""
	v.Address = ""
	v.Zip = ""
	poolAddressInfo.Put(v)
}
