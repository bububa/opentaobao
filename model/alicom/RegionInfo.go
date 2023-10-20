package alicom

import (
	"sync"
)

// RegionInfo 结构体
type RegionInfo struct {
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 省份id
	ProvId string `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// 省份
	Province string `json:"province,omitempty" xml:"province,omitempty"`
}

var poolRegionInfo = sync.Pool{
	New: func() any {
		return new(RegionInfo)
	},
}

// GetRegionInfo() 从对象池中获取RegionInfo
func GetRegionInfo() *RegionInfo {
	return poolRegionInfo.Get().(*RegionInfo)
}

// ReleaseRegionInfo 释放RegionInfo
func ReleaseRegionInfo(v *RegionInfo) {
	v.City = ""
	v.CityId = ""
	v.ProvId = ""
	v.Province = ""
	poolRegionInfo.Put(v)
}
