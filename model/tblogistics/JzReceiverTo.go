package tblogistics

import (
	"sync"
)

// JzReceiverTO 结构体
type JzReceiverTO struct {
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 街道
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 座机号
	TelePhone string `json:"tele_phone,omitempty" xml:"tele_phone,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
}

var poolJzReceiverTO = sync.Pool{
	New: func() any {
		return new(JzReceiverTO)
	},
}

// GetJzReceiverTO() 从对象池中获取JzReceiverTO
func GetJzReceiverTO() *JzReceiverTO {
	return poolJzReceiverTO.Get().(*JzReceiverTO)
}

// ReleaseJzReceiverTO 释放JzReceiverTO
func ReleaseJzReceiverTO(v *JzReceiverTO) {
	v.Address = ""
	v.City = ""
	v.ContactName = ""
	v.Country = ""
	v.District = ""
	v.MobilePhone = ""
	v.Province = ""
	v.Street = ""
	v.TelePhone = ""
	v.ZipCode = ""
	poolJzReceiverTO.Put(v)
}
