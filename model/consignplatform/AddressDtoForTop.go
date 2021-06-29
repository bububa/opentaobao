package consignplatform

// AddressDtoForTop 
type AddressDtoForTop struct {
    // 国家
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    // 省份
    ProvName   string `json:"prov_name,omitempty" xml:"prov_name,omitempty"`
    // 城市
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    // 区
    AreaName   string `json:"area_name,omitempty" xml:"area_name,omitempty"`
    // 街道
    TownName   string `json:"town_name,omitempty" xml:"town_name,omitempty"`
    // 详细地址
    AddressDetail   string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
}
