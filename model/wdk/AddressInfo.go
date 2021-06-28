package wdk

// AddressInfo 
/* model for simplify = false
type AddressInfo struct {

    // 地址类型
    
    AddressType   string `json:"address_type,omitempty"`
    

    // 城市
    
    City   string `json:"city,omitempty"`
    

    // 省份
    
    State   string `json:"state,omitempty"`
    

    // 国家
    
    Country   string `json:"country,omitempty"`
    

    // 详细地址
    
    Address   string `json:"address,omitempty"`
    

}
*/

// AddressInfo 
type AddressInfo struct {

    // 地址类型
    AddressType   string `json:"address_type,omitempty"`

    // 城市
    City   string `json:"city,omitempty"`

    // 省份
    State   string `json:"state,omitempty"`

    // 国家
    Country   string `json:"country,omitempty"`

    // 详细地址
    Address   string `json:"address,omitempty"`

}
