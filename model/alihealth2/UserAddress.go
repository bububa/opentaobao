package alihealth2

import (
	"sync"
)

// UserAddress 结构体
type UserAddress struct {
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 送货地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 买家姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

var poolUserAddress = sync.Pool{
	New: func() any {
		return new(UserAddress)
	},
}

// GetUserAddress() 从对象池中获取UserAddress
func GetUserAddress() *UserAddress {
	return poolUserAddress.Get().(*UserAddress)
}

// ReleaseUserAddress 释放UserAddress
func ReleaseUserAddress(v *UserAddress) {
	v.Phone = ""
	v.Address = ""
	v.Name = ""
	v.Mobile = ""
	poolUserAddress.Put(v)
}
