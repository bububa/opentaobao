package nropen

import (
	"sync"
)

// Divisioninfos 结构体
type Divisioninfos struct {
	// 扩展字段JSON字符串
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 国家
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 省份
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 城市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 区县
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 区域id
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}

var poolDivisioninfos = sync.Pool{
	New: func() any {
		return new(Divisioninfos)
	},
}

// GetDivisioninfos() 从对象池中获取Divisioninfos
func GetDivisioninfos() *Divisioninfos {
	return poolDivisioninfos.Get().(*Divisioninfos)
}

// ReleaseDivisioninfos 释放Divisioninfos
func ReleaseDivisioninfos(v *Divisioninfos) {
	v.Feature = ""
	v.CountryName = ""
	v.ProvinceName = ""
	v.CityName = ""
	v.DistrictName = ""
	v.DivisionId = 0
	poolDivisioninfos.Put(v)
}
