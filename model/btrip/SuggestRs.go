package btrip

// SuggestRs 
type SuggestRs struct {
    // 城市列表
    Cities   []CityVo `json:"cities,omitempty" xml:"cities>city_vo,omitempty"`
    // 是否为邻近城市
    Nearby   bool `json:"nearby,omitempty" xml:"nearby,omitempty"`
}
