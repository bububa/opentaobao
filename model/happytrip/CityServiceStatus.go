package happytrip

// CityServiceStatus 
type CityServiceStatus struct {

    // 供应商的城市id
    
    CityId   string `json:"city_id,omitempty" xml:"city_id,omitempty"`
    

    // 城市名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 支持的车型代码列表
    
    SupportServiceLevels   []string `json:"support_service_levels,omitempty" xml:"support_service_levels>string,omitempty"`
    

}
