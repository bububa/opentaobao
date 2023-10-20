package rail

import (
	"sync"
)

// RailDivisionRs 结构体
type RailDivisionRs struct {
	// 省份城市
	ProvName string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
	// 省份城市英文
	ProvNameEn string `json:"prov_name_en,omitempty" xml:"prov_name_en,omitempty"`
	// 国家名称英文
	CountryNameEn string `json:"country_name_en,omitempty" xml:"country_name_en,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 洲名称英文
	ContinentNameEn string `json:"continent_name_en,omitempty" xml:"continent_name_en,omitempty"`
	// 洲名称
	ContinentName string `json:"continent_name,omitempty" xml:"continent_name,omitempty"`
	// 城市三字码
	CityCode3 string `json:"city_code3,omitempty" xml:"city_code3,omitempty"`
	// 简拼
	Py string `json:"py,omitempty" xml:"py,omitempty"`
	// 拼音
	Pinyin string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
	// 时区
	Timezoneid string `json:"timezoneid,omitempty" xml:"timezoneid,omitempty"`
	// 国家二字码(ISO 3166-1 alpha-2)
	CountryCode2 string `json:"country_code2,omitempty" xml:"country_code2,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// 简称
	NameAbbr string `json:"name_abbr,omitempty" xml:"name_abbr,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 区名称英文
	DistrictNameEn string `json:"district_name_en,omitempty" xml:"district_name_en,omitempty"`
	// 区名称
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 城市名称英文
	CityNameEn string `json:"city_name_en,omitempty" xml:"city_name_en,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 省份id
	ProvId int64 `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// 国家id
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 洲id，id树层级展开
	ContinentId int64 `json:"continent_id,omitempty" xml:"continent_id,omitempty"`
	// 是否是首都
	CapitalFlag int64 `json:"capital_flag,omitempty" xml:"capital_flag,omitempty"`
	// 国内或者国际 0表示国内,1表示国际
	Abroad int64 `json:"abroad,omitempty" xml:"abroad,omitempty"`
	// 父级ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 层级，1洲，2是国家，3是省，4是市，5是区，6是街道/镇，7是村，8是逻辑行政区
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 区id
	DistrictId int64 `json:"district_id,omitempty" xml:"district_id,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 区域id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolRailDivisionRs = sync.Pool{
	New: func() any {
		return new(RailDivisionRs)
	},
}

// GetRailDivisionRs() 从对象池中获取RailDivisionRs
func GetRailDivisionRs() *RailDivisionRs {
	return poolRailDivisionRs.Get().(*RailDivisionRs)
}

// ReleaseRailDivisionRs 释放RailDivisionRs
func ReleaseRailDivisionRs(v *RailDivisionRs) {
	v.ProvName = ""
	v.ProvNameEn = ""
	v.CountryNameEn = ""
	v.CountryName = ""
	v.ContinentNameEn = ""
	v.ContinentName = ""
	v.CityCode3 = ""
	v.Py = ""
	v.Pinyin = ""
	v.Timezoneid = ""
	v.CountryCode2 = ""
	v.Latitude = ""
	v.Longitude = ""
	v.NameEn = ""
	v.NameAbbr = ""
	v.Name = ""
	v.DistrictNameEn = ""
	v.DistrictName = ""
	v.CityNameEn = ""
	v.CityName = ""
	v.ProvId = 0
	v.CountryId = 0
	v.ContinentId = 0
	v.CapitalFlag = 0
	v.Abroad = 0
	v.ParentId = 0
	v.Level = 0
	v.DistrictId = 0
	v.CityId = 0
	v.Id = 0
	poolRailDivisionRs.Put(v)
}
