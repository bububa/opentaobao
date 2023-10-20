package eleenterpriseordernew

import (
	"sync"
)

// AddressInfo 结构体
type AddressInfo struct {
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收货人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
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
	v.Address = ""
	v.Phone = ""
	v.Name = ""
	poolAddressInfo.Put(v)
}
