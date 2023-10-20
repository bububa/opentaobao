package travel

import (
	"sync"
)

// PontusTravelFreedomItemExt 结构体
type PontusTravelFreedomItemExt struct {
	// 其他资源信息
	OtherInfos []PontusTravelItemResourceInfo `json:"other_infos,omitempty" xml:"other_infos>pontus_travel_item_resource_info,omitempty"`
	// 自由行交通信息详细描述
	TrafficDesc string `json:"traffic_desc,omitempty" xml:"traffic_desc,omitempty"`
	// 回程交通信息
	BackTrafficInfo *PontusTravelItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
	// 去程交通信息
	GoTrafficInfo *PontusTravelItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
	// 酒店信息
	HotelInfos *PontusTravelItemHotelInfo `json:"hotel_infos,omitempty" xml:"hotel_infos,omitempty"`
	// 景点信息
	ScenicInfos *PontusTravelItemScenicInfo `json:"scenic_infos,omitempty" xml:"scenic_infos,omitempty"`
}

var poolPontusTravelFreedomItemExt = sync.Pool{
	New: func() any {
		return new(PontusTravelFreedomItemExt)
	},
}

// GetPontusTravelFreedomItemExt() 从对象池中获取PontusTravelFreedomItemExt
func GetPontusTravelFreedomItemExt() *PontusTravelFreedomItemExt {
	return poolPontusTravelFreedomItemExt.Get().(*PontusTravelFreedomItemExt)
}

// ReleasePontusTravelFreedomItemExt 释放PontusTravelFreedomItemExt
func ReleasePontusTravelFreedomItemExt(v *PontusTravelFreedomItemExt) {
	v.OtherInfos = v.OtherInfos[:0]
	v.TrafficDesc = ""
	v.BackTrafficInfo = nil
	v.GoTrafficInfo = nil
	v.HotelInfos = nil
	v.ScenicInfos = nil
	poolPontusTravelFreedomItemExt.Put(v)
}
