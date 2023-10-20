package tblogistics

import (
	"sync"
)

// ContactInfo 结构体
type ContactInfo struct {
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 固定电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区域
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 城镇
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// 详细地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 订单oaid，正向出库时下发
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
}

var poolContactInfo = sync.Pool{
	New: func() any {
		return new(ContactInfo)
	},
}

// GetContactInfo() 从对象池中获取ContactInfo
func GetContactInfo() *ContactInfo {
	return poolContactInfo.Get().(*ContactInfo)
}

// ReleaseContactInfo 释放ContactInfo
func ReleaseContactInfo(v *ContactInfo) {
	v.Name = ""
	v.Tel = ""
	v.Mobile = ""
	v.Province = ""
	v.City = ""
	v.Area = ""
	v.Town = ""
	v.DetailAddress = ""
	v.Oaid = ""
	poolContactInfo.Put(v)
}
