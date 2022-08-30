package axintrade

// ProductAreaInfoDto 结构体
type ProductAreaInfoDto struct {
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 地区名称
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 省份id
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 地区id
	DistrictId int64 `json:"district_id,omitempty" xml:"district_id,omitempty"`
}
