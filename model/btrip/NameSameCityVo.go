package btrip

// NameSameCityVo 结构体
type NameSameCityVo struct {
	// 来自城市
	FromCityList []CityVo `json:"from_city_list,omitempty" xml:"from_city_list>city_vo,omitempty"`
	// 到达城市
	ToCityList []CityVo `json:"to_city_list,omitempty" xml:"to_city_list>city_vo,omitempty"`
}
