package btrip

// AirportInfo 结构体
type AirportInfo struct {
	// 机场编码
	AirportCode string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	// 机场名称
	AirportName string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// 航站楼
	Terminal string `json:"terminal,omitempty" xml:"terminal,omitempty"`
	// 城市三字码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
}
