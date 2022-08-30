package alitripmerchant

// FacilityListVO 结构体
type FacilityListVO struct {
	// 设施Vo对象
	FacilityList []FacilityVO `json:"facility_list,omitempty" xml:"facility_list>facility_vo,omitempty"`
	// 设施类型名称
	FacilityName string `json:"facility_name,omitempty" xml:"facility_name,omitempty"`
}
