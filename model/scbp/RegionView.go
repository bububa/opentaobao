package scbp

// RegionView 
type RegionView struct {

    // 国家列表
    CountryList   []CountryView `json:"country_list,omitempty"`

    // 地区中文名
    RegionCnName   string `json:"region_cn_name,omitempty"`

}
