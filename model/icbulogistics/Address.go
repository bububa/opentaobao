package icbulogistics

// Address 
type Address struct {
    // 邮编
    Zip   string `json:"zip,omitempty" xml:"zip,omitempty"`
    // 国家
    Country   *Country `json:"country,omitempty" xml:"country,omitempty"`
    // 地址
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    // 省份
    Province   *Province `json:"province,omitempty" xml:"province,omitempty"`
    // 乡、镇名称
    Town   *Town `json:"town,omitempty" xml:"town,omitempty"`
    // 地址2
    Address2   string `json:"address2,omitempty" xml:"address2,omitempty"`
    // 城市
    City   *City `json:"city,omitempty" xml:"city,omitempty"`
    // 地区
    District   *District `json:"district,omitempty" xml:"district,omitempty"`
}
