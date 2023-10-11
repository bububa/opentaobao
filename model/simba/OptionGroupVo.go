package simba

// OptionGroupVo 结构体
type OptionGroupVo struct {
	// 标签选项
	OptionList []OptionVo `json:"option_list,omitempty" xml:"option_list>option_vo,omitempty"`
	// 标签分组名称
	OptionGroupName string `json:"option_group_name,omitempty" xml:"option_group_name,omitempty"`
}
