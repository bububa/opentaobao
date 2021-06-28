package logistic

// AddressDto 
type AddressDto struct {

    // 镇/街道
    
    TownName   string `json:"town_name,omitempty" xml:"town_name,omitempty"`
    

    // 详细地址
    
    AddressDetail   string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
    

    // 市
    
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    

    // 区
    
    AreaName   string `json:"area_name,omitempty" xml:"area_name,omitempty"`
    

    // 省
    
    ProvinceName   string `json:"province_name,omitempty" xml:"province_name,omitempty"`
    

    // 区级地址
    
    County   string `json:"county,omitempty" xml:"county,omitempty"`
    

    // 省级地址
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 街道地址
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

    // 详细地址
    
    DetailAddress   string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
    

    // 市级地址
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 国家地址
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

}
