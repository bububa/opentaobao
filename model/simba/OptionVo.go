package simba

// OptionVo 结构体
type OptionVo struct {
	// 选项名称
	OptionName string `json:"option_name,omitempty" xml:"option_name,omitempty"`
	// 选项值
	OptionValue string `json:"option_value,omitempty" xml:"option_value,omitempty"`
	// 选项描述
	OptionDesc string `json:"option_desc,omitempty" xml:"option_desc,omitempty"`
	// 标签分组名称
	OptionGroupName string `json:"option_group_name,omitempty" xml:"option_group_name,omitempty"`
	// 额外的属性参数
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 选项id，(用于关键词推广静态标签生成人群)如基础属性人群标签组合
	TagId int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
	// 透传属性
	LabelOptionProperties *LabelOptionProperties `json:"label_option_properties,omitempty" xml:"label_option_properties,omitempty"`
}
