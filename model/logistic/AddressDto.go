package logistic

// AddressDto 
type AddressDto struct {

    // 镇/街道
    TownName   string `json:"town_name,omitempty"`

    // 详细地址
    AddressDetail   string `json:"address_detail,omitempty"`

    // 市
    CityName   string `json:"city_name,omitempty"`

    // 区
    AreaName   string `json:"area_name,omitempty"`

    // 省
    ProvinceName   string `json:"province_name,omitempty"`

    // 区级地址
    County   string `json:"county,omitempty"`

    // 省级地址
    Province   string `json:"province,omitempty"`

    // 街道地址
    Town   string `json:"town,omitempty"`

    // 详细地址
    DetailAddress   string `json:"detail_address,omitempty"`

    // 市级地址
    City   string `json:"city,omitempty"`

    // 国家地址
    Country   string `json:"country,omitempty"`

}
