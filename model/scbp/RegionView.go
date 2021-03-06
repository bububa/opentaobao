package scbp

// RegionView 结构体
type RegionView struct {
	// 国家列表
	CountryList []CountryView `json:"country_list,omitempty" xml:"country_list>country_view,omitempty"`
	// 地区中文名
	RegionCnName string `json:"region_cn_name,omitempty" xml:"region_cn_name,omitempty"`
}
