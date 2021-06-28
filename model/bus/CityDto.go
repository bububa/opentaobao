package bus

// CityDto 
/* model for simplify = false
type CityDto struct {

    // 城市编码
    
    CityCode   string `json:"city_code,omitempty"`
    

    // 拼音
    
    CityFullPy   string `json:"city_full_py,omitempty"`
    

    // 城市名字
    
    CityName   string `json:"city_name,omitempty"`
    

    // 简拼
    
    CityShortPy   string `json:"city_short_py,omitempty"`
    

    // 预售期
    
    PreDay   int64 `json:"pre_day,omitempty"`
    

    // 省份名称
    
    ProvinceName   string `json:"province_name,omitempty"`
    

}
*/

// CityDto 
type CityDto struct {

    // 城市编码
    CityCode   string `json:"city_code,omitempty"`

    // 拼音
    CityFullPy   string `json:"city_full_py,omitempty"`

    // 城市名字
    CityName   string `json:"city_name,omitempty"`

    // 简拼
    CityShortPy   string `json:"city_short_py,omitempty"`

    // 预售期
    PreDay   int64 `json:"pre_day,omitempty"`

    // 省份名称
    ProvinceName   string `json:"province_name,omitempty"`

}
