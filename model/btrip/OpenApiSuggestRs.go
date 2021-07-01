package btrip

// OpenApiSuggestRS 
type OpenApiSuggestRS struct {
    // 城市列表
    Cities   []CityVo `json:"cities,omitempty" xml:"cities>city_vo,omitempty"`
}
