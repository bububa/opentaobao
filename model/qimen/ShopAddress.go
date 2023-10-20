package qimen

import (
	"sync"
)

// ShopAddress 结构体
type ShopAddress struct {
	// 邮编, string (50)
	ZipCode string `json:"zipCode,omitempty" xml:"zipCode,omitempty"`
	// 省份, string (50)
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市, string (50)
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区域, string (50)
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 村镇, string (50)
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址, string (200)
	DetailAddress string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
}

var poolShopAddress = sync.Pool{
	New: func() any {
		return new(ShopAddress)
	},
}

// GetShopAddress() 从对象池中获取ShopAddress
func GetShopAddress() *ShopAddress {
	return poolShopAddress.Get().(*ShopAddress)
}

// ReleaseShopAddress 释放ShopAddress
func ReleaseShopAddress(v *ShopAddress) {
	v.ZipCode = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	poolShopAddress.Put(v)
}
