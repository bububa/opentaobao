package travel

import (
	"sync"
)

// PontusTravelGatherPlaceInfo 结构体
type PontusTravelGatherPlaceInfo struct {
	// 地点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 集合地经纬度，英文逗号分隔。经度在前，纬度在后，英文逗号分隔最多支持6位小数，如120.111222,30.111222
	Poi string `json:"poi,omitempty" xml:"poi,omitempty"`
	// POI来源，AMAP/GOOGLE。境内为高德（AMAP） 境外为GOOGLE
	PoiResource string `json:"poi_resource,omitempty" xml:"poi_resource,omitempty"`
}

var poolPontusTravelGatherPlaceInfo = sync.Pool{
	New: func() any {
		return new(PontusTravelGatherPlaceInfo)
	},
}

// GetPontusTravelGatherPlaceInfo() 从对象池中获取PontusTravelGatherPlaceInfo
func GetPontusTravelGatherPlaceInfo() *PontusTravelGatherPlaceInfo {
	return poolPontusTravelGatherPlaceInfo.Get().(*PontusTravelGatherPlaceInfo)
}

// ReleasePontusTravelGatherPlaceInfo 释放PontusTravelGatherPlaceInfo
func ReleasePontusTravelGatherPlaceInfo(v *PontusTravelGatherPlaceInfo) {
	v.Name = ""
	v.Poi = ""
	v.PoiResource = ""
	poolPontusTravelGatherPlaceInfo.Put(v)
}
