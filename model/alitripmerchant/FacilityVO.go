package alitripmerchant

// FacilityVO 结构体
type FacilityVO struct {
	// 图片icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 设施名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 概要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 设施Id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
