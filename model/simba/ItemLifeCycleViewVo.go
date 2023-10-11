package simba

// ItemLifeCycleViewVo 结构体
type ItemLifeCycleViewVo struct {
	// 周期文案
	LifeCycleDesc string `json:"life_cycle_desc,omitempty" xml:"life_cycle_desc,omitempty"`
	// 文案颜色
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// 周期提示文案
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
}
