package alihealthoutflow

import (
	"sync"
)

// AddressVo 结构体
type AddressVo struct {
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
}

var poolAddressVo = sync.Pool{
	New: func() any {
		return new(AddressVo)
	},
}

// GetAddressVo() 从对象池中获取AddressVo
func GetAddressVo() *AddressVo {
	return poolAddressVo.Get().(*AddressVo)
}

// ReleaseAddressVo 释放AddressVo
func ReleaseAddressVo(v *AddressVo) {
	v.Province = ""
	v.City = ""
	poolAddressVo.Put(v)
}
