package ascp

import (
	"sync"
)

// AddressNames 结构体
type AddressNames struct {
	// 浙江省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
}

var poolAddressNames = sync.Pool{
	New: func() any {
		return new(AddressNames)
	},
}

// GetAddressNames() 从对象池中获取AddressNames
func GetAddressNames() *AddressNames {
	return poolAddressNames.Get().(*AddressNames)
}

// ReleaseAddressNames 释放AddressNames
func ReleaseAddressNames(v *AddressNames) {
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	poolAddressNames.Put(v)
}
