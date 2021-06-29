package nropen

// Divisioninfos 
type Divisioninfos struct {

    // 区域id
    
    DivisionId   int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
    

    // 扩展字段JSON字符串
    
    Feature   string `json:"feature,omitempty" xml:"feature,omitempty"`
    

    // 国家
    
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    

    // 省份
    
    ProvinceName   string `json:"province_name,omitempty" xml:"province_name,omitempty"`
    

    // 城市
    
    CityName   string `json:"city_name,omitempty" xml:"city_name,omitempty"`
    

    // 区县
    
    DistrictName   string `json:"district_name,omitempty" xml:"district_name,omitempty"`
    

}
