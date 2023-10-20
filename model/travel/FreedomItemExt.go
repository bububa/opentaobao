package travel

import (
	"sync"
)

// FreedomItemExt 结构体
type FreedomItemExt struct {
	// 其他资源信息
	OtherInfos []ItemResourceInfo `json:"other_infos,omitempty" xml:"other_infos>item_resource_info,omitempty"`
	// 自由行交通信息详细描述
	TrafficDesc string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
	// 回程交通信息
	BackTrafficInfo *ItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
	// 去程交通信息
	GoTrafficInfo *ItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
	// 酒店信息
	HotelInfos *ItemHotelInfo `json:"hotel_infos,omitempty" xml:"hotel_infos,omitempty"`
	// 景点信息
	ScenicInfos *ItemScenicInfo `json:"scenic_infos,omitempty" xml:"scenic_infos,omitempty"`
}

var poolFreedomItemExt = sync.Pool{
	New: func() any {
		return new(FreedomItemExt)
	},
}

// GetFreedomItemExt() 从对象池中获取FreedomItemExt
func GetFreedomItemExt() *FreedomItemExt {
	return poolFreedomItemExt.Get().(*FreedomItemExt)
}

// ReleaseFreedomItemExt 释放FreedomItemExt
func ReleaseFreedomItemExt(v *FreedomItemExt) {
	v.OtherInfos = v.OtherInfos[:0]
	v.TrafficDesc = ""
	v.BackTrafficInfo = nil
	v.GoTrafficInfo = nil
	v.HotelInfos = nil
	v.ScenicInfos = nil
	poolFreedomItemExt.Put(v)
}
