package axindata

// FscPoiApplyApiDto 结构体
type FscPoiApplyApiDto struct {
	// POI外部编号（供应商侧编号）
	PoiOuterId string `json:"poi_outer_id,omitempty" xml:"poi_outer_id,omitempty"`
	// POI名称
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// POI英文名称
	PoiNameEn string `json:"poi_name_en,omitempty" xml:"poi_name_en,omitempty"`
	// POI详情说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 城市ID
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 省份ID
	ProvinceId string `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 国家ID
	CountryId string `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
}
