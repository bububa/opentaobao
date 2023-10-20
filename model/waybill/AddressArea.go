package waybill

import (
	"sync"
)

// AddressArea 结构体
type AddressArea struct {
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
}

var poolAddressArea = sync.Pool{
	New: func() any {
		return new(AddressArea)
	},
}

// GetAddressArea() 从对象池中获取AddressArea
func GetAddressArea() *AddressArea {
	return poolAddressArea.Get().(*AddressArea)
}

// ReleaseAddressArea 释放AddressArea
func ReleaseAddressArea(v *AddressArea) {
	v.City = ""
	v.Province = ""
	poolAddressArea.Put(v)
}
