package flightuppc

// AlitripFlightBasicDataCityQueryAllData 结构体
type AlitripFlightBasicDataCityQueryAllData struct {
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 城市三字码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市通航状态
	NavigableStatus int64 `json:"navigable_status,omitempty" xml:"navigable_status,omitempty"`
	// 搜索热度
	SearchHeat int64 `json:"search_heat,omitempty" xml:"search_heat,omitempty"`
	// 航空大区
	AirZone string `json:"air_zone,omitempty" xml:"air_zone,omitempty"`
	// 城市对应国家
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 城市简拼
	CityJp string `json:"city_jp,omitempty" xml:"city_jp,omitempty"`
	// 夏日时区区间
	TimeZoneSummerPeriod string `json:"time_zone_summer_period,omitempty" xml:"time_zone_summer_period,omitempty"`
	// 夏日时区
	TimeZoneSummer string `json:"time_zone_summer,omitempty" xml:"time_zone_summer,omitempty"`
	// 标准时区
	TimeZoneStandard string `json:"time_zone_standard,omitempty" xml:"time_zone_standard,omitempty"`
	// 大洲
	Continent string `json:"continent,omitempty" xml:"continent,omitempty"`
	// 省/州名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 国家二字码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 省/州二字码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 首字母
	CityCapital string `json:"city_capital,omitempty" xml:"city_capital,omitempty"`
	// 英文名称
	EnglishName string `json:"english_name,omitempty" xml:"english_name,omitempty"`
	// 城市全拼
	CityPy string `json:"city_py,omitempty" xml:"city_py,omitempty"`
	// 次区代码
	IataCode string `json:"iata_code,omitempty" xml:"iata_code,omitempty"`
	// 次区
	SecdZone string `json:"secd_zone,omitempty" xml:"secd_zone,omitempty"`
	// OAG来源的夏令时
	OagSummerPeriod string `json:"oag_summer_period,omitempty" xml:"oag_summer_period,omitempty"`
	// 小次区
	MiniIataCode string `json:"mini_iata_code,omitempty" xml:"mini_iata_code,omitempty"`
}
