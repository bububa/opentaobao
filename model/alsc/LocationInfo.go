package alsc

import (
	"sync"
)

// LocationInfo 结构体
type LocationInfo struct {
	// 市区域码
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区县区域码
	District string `json:"district,omitempty" xml:"district,omitempty"`
	//  纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 位置类型(默认实时)
	LocationType string `json:"location_type,omitempty" xml:"location_type,omitempty"`
	// poiId
	PoiId string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	//  poiName
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// 省区域码
	Province string `json:"province,omitempty" xml:"province,omitempty"`
}

var poolLocationInfo = sync.Pool{
	New: func() any {
		return new(LocationInfo)
	},
}

// GetLocationInfo() 从对象池中获取LocationInfo
func GetLocationInfo() *LocationInfo {
	return poolLocationInfo.Get().(*LocationInfo)
}

// ReleaseLocationInfo 释放LocationInfo
func ReleaseLocationInfo(v *LocationInfo) {
	v.City = ""
	v.District = ""
	v.Lat = ""
	v.Lng = ""
	v.LocationType = ""
	v.PoiId = ""
	v.PoiName = ""
	v.Province = ""
	poolLocationInfo.Put(v)
}
