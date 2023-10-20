package ascp

import (
	"sync"
)

// AddressName 结构体
type AddressName struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}

var poolAddressName = sync.Pool{
	New: func() any {
		return new(AddressName)
	},
}

// GetAddressName() 从对象池中获取AddressName
func GetAddressName() *AddressName {
	return poolAddressName.Get().(*AddressName)
}

// ReleaseAddressName 释放AddressName
func ReleaseAddressName(v *AddressName) {
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	poolAddressName.Put(v)
}
