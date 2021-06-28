package wdk

// AddressInfo 
type AddressInfo struct {

    // 地址类型
    
    AddressType   string `json:"address_type,omitempty" xml:"address_type,omitempty"`
    

    // 城市
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 省份
    
    State   string `json:"state,omitempty" xml:"state,omitempty"`
    

    // 国家
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

    // 详细地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

}
