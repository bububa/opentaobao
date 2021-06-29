package icbulogistics

// AddressQueryDTO 
type AddressQueryDTO struct {
    // 省ID
    ProvinceId   int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
    // 城市id
    CityId   int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
    // 是否包含子节点
    WithChildren   bool `json:"with_children,omitempty" xml:"with_children,omitempty"`
    // 国家code
    CountryCode   string `json:"country_code,omitempty" xml:"country_code,omitempty"`
    // 查询关键词
    SearchText   string `json:"search_text,omitempty" xml:"search_text,omitempty"`
}
