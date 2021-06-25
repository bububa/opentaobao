package wlbimports

// ReciverAddressDo 
type ReciverAddressDo struct {

    // 详细地址
    DetailAddress   string `json:"detail_address,omitempty"`

    // 街道
    Street   string `json:"street,omitempty"`

    // 省级别
    Province   string `json:"province,omitempty"`

    // 区、县级别
    District   string `json:"district,omitempty"`

    // 国级别
    Country   string `json:"country,omitempty"`

    // 市级别
    City   string `json:"city,omitempty"`

}