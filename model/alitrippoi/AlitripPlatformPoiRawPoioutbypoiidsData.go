package alitrippoi

import (
	"sync"
)

// AlitripPlatformPoiRawPoioutbypoiidsData 结构体
type AlitripPlatformPoiRawPoioutbypoiidsData struct {
	// poiId
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 行政区划树名
	DivisionTreeName string `json:"division_tree_name,omitempty" xml:"division_tree_name,omitempty"`
	// 行政区划树
	DivisionTreeId string `json:"division_tree_id,omitempty" xml:"division_tree_id,omitempty"`
	// 开放时间
	OpenTime string `json:"open_time,omitempty" xml:"open_time,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lon string `json:"lon,omitempty" xml:"lon,omitempty"`
	// 电话
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// poi英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// poi名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类目
	FirstCategory int64 `json:"first_category,omitempty" xml:"first_category,omitempty"`
	// poiId
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

var poolAlitripPlatformPoiRawPoioutbypoiidsData = sync.Pool{
	New: func() any {
		return new(AlitripPlatformPoiRawPoioutbypoiidsData)
	},
}

// GetAlitripPlatformPoiRawPoioutbypoiidsData() 从对象池中获取AlitripPlatformPoiRawPoioutbypoiidsData
func GetAlitripPlatformPoiRawPoioutbypoiidsData() *AlitripPlatformPoiRawPoioutbypoiidsData {
	return poolAlitripPlatformPoiRawPoioutbypoiidsData.Get().(*AlitripPlatformPoiRawPoioutbypoiidsData)
}

// ReleaseAlitripPlatformPoiRawPoioutbypoiidsData 释放AlitripPlatformPoiRawPoioutbypoiidsData
func ReleaseAlitripPlatformPoiRawPoioutbypoiidsData(v *AlitripPlatformPoiRawPoioutbypoiidsData) {
	v.SourceId = ""
	v.DivisionTreeName = ""
	v.DivisionTreeId = ""
	v.OpenTime = ""
	v.Lat = ""
	v.Lon = ""
	v.Telephone = ""
	v.Address = ""
	v.Desc = ""
	v.NameEn = ""
	v.Name = ""
	v.FirstCategory = 0
	v.PoiId = 0
	poolAlitripPlatformPoiRawPoioutbypoiidsData.Put(v)
}
