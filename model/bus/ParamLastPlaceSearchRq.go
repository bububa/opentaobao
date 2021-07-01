package bus

// ParamLastPlaceSearchRq 结构体
type ParamLastPlaceSearchRq struct {
	// 城市编码
	DepCityCode string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// 北京市
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
}
