package omniorder

import (
	"sync"
)

// JzReceiverDto 结构体
type JzReceiverDto struct {
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 收货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 街道
	Street string `json:"street,omitempty" xml:"street,omitempty"`
	// 区
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 座机号
	TelePhone string `json:"tele_phone,omitempty" xml:"tele_phone,omitempty"`
}

var poolJzReceiverDto = sync.Pool{
	New: func() any {
		return new(JzReceiverDto)
	},
}

// GetJzReceiverDto() 从对象池中获取JzReceiverDto
func GetJzReceiverDto() *JzReceiverDto {
	return poolJzReceiverDto.Get().(*JzReceiverDto)
}

// ReleaseJzReceiverDto 释放JzReceiverDto
func ReleaseJzReceiverDto(v *JzReceiverDto) {
	v.Country = ""
	v.ZipCode = ""
	v.Address = ""
	v.MobilePhone = ""
	v.Province = ""
	v.City = ""
	v.ContactName = ""
	v.Street = ""
	v.District = ""
	v.TelePhone = ""
	poolJzReceiverDto.Put(v)
}
