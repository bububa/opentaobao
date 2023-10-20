package icbulogistics

import (
	"sync"
)

// Address 结构体
type Address struct {
	// 邮编
	Zip string `json:"zip,omitempty" xml:"zip,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 地址2
	Address2 string `json:"address2,omitempty" xml:"address2,omitempty"`
	// 国家
	Country *Country `json:"country,omitempty" xml:"country,omitempty"`
	// 乡、镇名称
	Town *Town `json:"town,omitempty" xml:"town,omitempty"`
	// 省份
	Province *Province `json:"province,omitempty" xml:"province,omitempty"`
	// 城市
	City *City `json:"city,omitempty" xml:"city,omitempty"`
	// 地区
	District *District `json:"district,omitempty" xml:"district,omitempty"`
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
	v.Zip = ""
	v.Address = ""
	v.Address2 = ""
	v.Country = nil
	v.Town = nil
	v.Province = nil
	v.City = nil
	v.District = nil
	poolAddress.Put(v)
}
