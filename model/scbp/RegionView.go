package scbp

// RegionView 
/* model for simplify = false
type RegionView struct {

    // 国家列表
    
    CountryList  struct {
        CountryView  []CountryView `json:"country_view,omitempty"`
    } `json:"country_list,omitempty"`
    

    // 地区中文名
    
    RegionCnName   string `json:"region_cn_name,omitempty"`
    

}
*/

// RegionView 
type RegionView struct {

    // 国家列表
    CountryList   []CountryView `json:"country_list,omitempty"`

    // 地区中文名
    RegionCnName   string `json:"region_cn_name,omitempty"`

}
