package tuanhotel

import (
	"sync"
)

// AppointPhoneVoList 结构体
type AppointPhoneVoList struct {
	// ext
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 地区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 手机
	Mobil string `json:"mobil,omitempty" xml:"mobil,omitempty"`
	// 国家
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 电话
	Fix string `json:"fix,omitempty" xml:"fix,omitempty"`
	// p400
	P400 string `json:"p400,omitempty" xml:"p400,omitempty"`
	// 固话或者移动电话
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolAppointPhoneVoList = sync.Pool{
	New: func() any {
		return new(AppointPhoneVoList)
	},
}

// GetAppointPhoneVoList() 从对象池中获取AppointPhoneVoList
func GetAppointPhoneVoList() *AppointPhoneVoList {
	return poolAppointPhoneVoList.Get().(*AppointPhoneVoList)
}

// ReleaseAppointPhoneVoList 释放AppointPhoneVoList
func ReleaseAppointPhoneVoList(v *AppointPhoneVoList) {
	v.Ext = ""
	v.Area = ""
	v.Mobil = ""
	v.Country = ""
	v.Fix = ""
	v.P400 = ""
	v.Type = ""
	poolAppointPhoneVoList.Put(v)
}
