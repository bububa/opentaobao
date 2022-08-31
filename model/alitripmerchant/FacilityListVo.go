package alitripmerchant

// FacilityListVo 结构体
type FacilityListVo struct {
	// 设施集合
	FacilityList []FacilityVo `json:"facility_list,omitempty" xml:"facility_list>facility_vo,omitempty"`
	// 设施分组名
	FacilityName string `json:"facility_name,omitempty" xml:"facility_name,omitempty"`
}
