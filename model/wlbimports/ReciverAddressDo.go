package wlbimports

import (
	"sync"
)

// ReciverAddressDo 结构体
type ReciverAddressDo struct {
	// 详细地址，只需填写买家具体的收货地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 街道信息
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 区县地址信息，city和district不能同时为空
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 市地址信息，city和district不能同时为空
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省地址信息
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 国家地址信息
	Country string `json:"country,omitempty" xml:"country,omitempty"`
}

var poolReciverAddressDo = sync.Pool{
	New: func() any {
		return new(ReciverAddressDo)
	},
}

// GetReciverAddressDo() 从对象池中获取ReciverAddressDo
func GetReciverAddressDo() *ReciverAddressDo {
	return poolReciverAddressDo.Get().(*ReciverAddressDo)
}

// ReleaseReciverAddressDo 释放ReciverAddressDo
func ReleaseReciverAddressDo(v *ReciverAddressDo) {
	v.DetailAddress = ""
	v.Street = ""
	v.District = ""
	v.City = ""
	v.Province = ""
	v.Country = ""
	poolReciverAddressDo.Put(v)
}
