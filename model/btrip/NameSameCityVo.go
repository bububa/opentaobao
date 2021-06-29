package btrip

// NameSameCityVo 
type NameSameCityVo struct {

    // 来自城市
    
    FromCityList   []CityVo `json:"from_city_list,omitempty" xml:"from_city_list,omitempty"`
    

    // 到达城市
    
    ToCityList   []CityVo `json:"to_city_list,omitempty" xml:"to_city_list,omitempty"`
    

}
