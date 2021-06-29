package perfect

// PerfectItemLogisticsInfoDto 
type PerfectItemLogisticsInfoDto struct {
    // 城市编码,默认北京
    CityCode   string `json:"city_code,omitempty" xml:"city_code,omitempty"`
    // 运费模板ID
    PostageTemplateCode   string `json:"postage_template_code,omitempty" xml:"postage_template_code,omitempty"`
    // 省份编码,默认北京
    ProvinceCode   string `json:"province_code,omitempty" xml:"province_code,omitempty"`
    // 门店库ID
    StoreGroupCode   string `json:"store_group_code,omitempty" xml:"store_group_code,omitempty"`
}
