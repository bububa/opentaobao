package alitripmerchant

// FacilityVO 结构体
type FacilityVO struct {
	// 文案表达
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 设施图标URL
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 设施名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 补充说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 设施Id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
