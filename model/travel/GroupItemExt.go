package travel

import (
	"sync"
)

// GroupItemExt 结构体
type GroupItemExt struct {
	// 集合地信息
	GatherPlaces []GatherPlaceInfo `json:"gather_places,omitempty" xml:"gather_places>gather_place_info,omitempty"`
	// 必填，回程交通信息
	BackTrafficInfo *ItemTrafficInfo `json:"back_traffic_info,omitempty" xml:"back_traffic_info,omitempty"`
	// 必填，去程交通信息
	GoTrafficInfo *ItemTrafficInfo `json:"go_traffic_info,omitempty" xml:"go_traffic_info,omitempty"`
	// 必填，线路类型，0 为目的地参团 	1为出发地参团
	RouteType int64 `json:"route_type,omitempty" xml:"route_type,omitempty"`
	// 是否支持电子合同，默认不支持
	Electronic bool `json:"electronic,omitempty" xml:"electronic,omitempty"`
}

var poolGroupItemExt = sync.Pool{
	New: func() any {
		return new(GroupItemExt)
	},
}

// GetGroupItemExt() 从对象池中获取GroupItemExt
func GetGroupItemExt() *GroupItemExt {
	return poolGroupItemExt.Get().(*GroupItemExt)
}

// ReleaseGroupItemExt 释放GroupItemExt
func ReleaseGroupItemExt(v *GroupItemExt) {
	v.GatherPlaces = v.GatherPlaces[:0]
	v.BackTrafficInfo = nil
	v.GoTrafficInfo = nil
	v.RouteType = 0
	v.Electronic = false
	poolGroupItemExt.Put(v)
}
