package btrip

// OpenApiSuggestRs 结构体
type OpenApiSuggestRs struct {
	// 城市列表
	Cities []CityVo `json:"cities,omitempty" xml:"cities>city_vo,omitempty"`
}
