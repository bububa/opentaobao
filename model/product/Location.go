package product

// Location 
type Location struct {

    // 邮政编码
    Zip   string `json:"zip,omitempty"`

    // 详细地址，最大256个字节（128个中文）
    Address   string `json:"address,omitempty"`

    // 所在城市（中文名称）
    City   string `json:"city,omitempty"`

    // 所在省份（中文名称）
    State   string `json:"state,omitempty"`

    // 国家名称
    Country   string `json:"country,omitempty"`

    // 区/县（只适用于物流API）
    District   string `json:"district,omitempty"`

}
