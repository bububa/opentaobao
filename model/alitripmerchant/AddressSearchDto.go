package alitripmerchant

// AddressSearchDto 
type AddressSearchDto struct {

    // 0国内1国外
    
    Domestic   int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
    

    // 城市图片url
    
    CityUrl   string `json:"city_url,omitempty" xml:"city_url,omitempty"`
    

    // 城市拼音首字母
    
    CityPyHead   string `json:"city_py_head,omitempty" xml:"city_py_head,omitempty"`
    

    // 国家编码
    
    CountryCode   int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
    

    // 城市编码
    
    CityCode   int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
    

    // 城市
    
    CityCn   string `json:"city_cn,omitempty" xml:"city_cn,omitempty"`
    

    // 国家
    
    CountryCn   string `json:"country_cn,omitempty" xml:"country_cn,omitempty"`
    

}
