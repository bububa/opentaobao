package alitrippoi

import (
	"sync"
)

// StructureServiceInfo 结构体
type StructureServiceInfo struct {
	// 是否提供酒精饮料
	Alcohol bool `json:"alcohol,omitempty" xml:"alcohol,omitempty"`
	// 是否有停车位
	Parking bool `json:"parking,omitempty" xml:"parking,omitempty"`
	// 是否有wifi
	Wifi bool `json:"wifi,omitempty" xml:"wifi,omitempty"`
	// 是否接受预约
	Booking bool `json:"booking,omitempty" xml:"booking,omitempty"`
	// 是否允许自带杯
	Byo bool `json:"byo,omitempty" xml:"byo,omitempty"`
	// 是否有包厢
	Box bool `json:"box,omitempty" xml:"box,omitempty"`
	// 是否提供中文服务
	ChineseSvc bool `json:"chinese_svc,omitempty" xml:"chinese_svc,omitempty"`
	// 是否支持外带
	Takeout bool `json:"takeout,omitempty" xml:"takeout,omitempty"`
	// 是否需要小费
	Tips bool `json:"tips,omitempty" xml:"tips,omitempty"`
	// 是否支持电话预定
	TelRsvt bool `json:"tel_rsvt,omitempty" xml:"tel_rsvt,omitempty"`
}

var poolStructureServiceInfo = sync.Pool{
	New: func() any {
		return new(StructureServiceInfo)
	},
}

// GetStructureServiceInfo() 从对象池中获取StructureServiceInfo
func GetStructureServiceInfo() *StructureServiceInfo {
	return poolStructureServiceInfo.Get().(*StructureServiceInfo)
}

// ReleaseStructureServiceInfo 释放StructureServiceInfo
func ReleaseStructureServiceInfo(v *StructureServiceInfo) {
	v.Alcohol = false
	v.Parking = false
	v.Wifi = false
	v.Booking = false
	v.Byo = false
	v.Box = false
	v.ChineseSvc = false
	v.Takeout = false
	v.Tips = false
	v.TelRsvt = false
	poolStructureServiceInfo.Put(v)
}
