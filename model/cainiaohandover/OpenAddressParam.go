package cainiaohandover

// OpenAddressParam 
type OpenAddressParam struct {
    // 最小区划地址库ID
    DivisionId   int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
    // 邮编
    ZipCode   string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
    // 国家名称
    CountryName   string `json:"country_name,omitempty" xml:"country_name,omitempty"`
    // 省份
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    // 城市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 区
    District   string `json:"district,omitempty" xml:"district,omitempty"`
    // 街道
    Street   string `json:"street,omitempty" xml:"street,omitempty"`
    // 详细地址
    DetailAddress   string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
    // 国家二字码，CN：中国、RU：俄罗斯
    CountryCode   string `json:"country_code,omitempty" xml:"country_code,omitempty"`
}
