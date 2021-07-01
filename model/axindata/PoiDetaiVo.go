package axindata

// PoiDetaiVo 结构体
type PoiDetaiVo struct {
	// poi描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 国家名称
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 国家id
	CountryId string `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 省份
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 省份id
	ProvinceId string `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 城市
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// poi英文名称
	PoiNameEn string `json:"poi_name_en,omitempty" xml:"poi_name_en,omitempty"`
	// poi名称
	PoiName string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	// poiId
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}
