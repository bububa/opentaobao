package waybill

import (
	"sync"
)

// Address 结构体
type Address struct {
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 详细地址
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
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
	v.City = ""
	v.Detail = ""
	v.District = ""
	v.Province = ""
	v.Town = ""
	poolAddress.Put(v)
}
