package axintrade

// AxinPayRegisterAddress 结构体
type AxinPayRegisterAddress struct {
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 区县编码
	DistrictCode string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 城市code
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 省code
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 国家code
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
}
