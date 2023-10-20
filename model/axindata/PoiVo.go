package axindata

import (
	"sync"
)

// PoiVo 结构体
type PoiVo struct {
	// poi英文名
	PoiNameEn string `json:"poi_name_en,omitempty" xml:"poi_name_en,omitempty"`
	// poi名称
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 国家id
	CountryId string `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 省份id
	ProvinceId string `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// poiId
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

var poolPoiVo = sync.Pool{
	New: func() any {
		return new(PoiVo)
	},
}

// GetPoiVo() 从对象池中获取PoiVo
func GetPoiVo() *PoiVo {
	return poolPoiVo.Get().(*PoiVo)
}

// ReleasePoiVo 释放PoiVo
func ReleasePoiVo(v *PoiVo) {
	v.PoiNameEn = ""
	v.PoiName = ""
	v.CityName = ""
	v.CityId = ""
	v.CountryName = ""
	v.CountryId = ""
	v.ProvinceName = ""
	v.ProvinceId = ""
	v.PoiId = 0
	poolPoiVo.Put(v)
}
