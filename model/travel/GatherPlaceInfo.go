package travel

import (
	"sync"
)

// GatherPlaceInfo 结构体
type GatherPlaceInfo struct {
	// 集合地点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 集合地经纬度，英文逗号分隔。经度在前，纬度在后，英文逗号分隔最多支持6位小数，如120.111222,30.111222
	Poi string `json:"poi,omitempty" xml:"poi,omitempty"`
	// AMAP/GOOGLE/OTHERS。高德（AMAP），GOOGLE，其他（OTHERS）
	PoiResource string `json:"poi_resource,omitempty" xml:"poi_resource,omitempty"`
}

var poolGatherPlaceInfo = sync.Pool{
	New: func() any {
		return new(GatherPlaceInfo)
	},
}

// GetGatherPlaceInfo() 从对象池中获取GatherPlaceInfo
func GetGatherPlaceInfo() *GatherPlaceInfo {
	return poolGatherPlaceInfo.Get().(*GatherPlaceInfo)
}

// ReleaseGatherPlaceInfo 释放GatherPlaceInfo
func ReleaseGatherPlaceInfo(v *GatherPlaceInfo) {
	v.Name = ""
	v.Poi = ""
	v.PoiResource = ""
	poolGatherPlaceInfo.Put(v)
}
