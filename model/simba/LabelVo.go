package simba

// LabelVo 结构体
type LabelVo struct {
	// 标签选项分组
	OptionGroupList []OptionGroupVo `json:"option_group_list,omitempty" xml:"option_group_list>option_group_vo,omitempty"`
	// 标签选项
	OptionList []OptionVo `json:"option_list,omitempty" xml:"option_list>option_vo,omitempty"`
	// 标签tag
	ShowTagList []ShowTagVo `json:"show_tag_list,omitempty" xml:"show_tag_list>show_tag_vo,omitempty"`
	// 定向标签名称
	LabelName string `json:"label_name,omitempty" xml:"label_name,omitempty"`
	// 定向标签描述
	LabelDesc string `json:"label_desc,omitempty" xml:"label_desc,omitempty"`
	// 定向标签值
	LabelValue string `json:"label_value,omitempty" xml:"label_value,omitempty"`
	// 出价维度，OPTION:标签出价，多个子人群同时出价、LABEL:选项出价，以子人群为出价最小单位
	PriceDimension string `json:"price_dimension,omitempty" xml:"price_dimension,omitempty"`
	// 人群类型，所属类型
	LabelGroupName string `json:"label_group_name,omitempty" xml:"label_group_name,omitempty"`
	// 定向标签类型
	TargetType int64 `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 定向标签id
	LabelId int64 `json:"label_id,omitempty" xml:"label_id,omitempty"`
	// 透传属性
	LabelOptionProperties *LabelOptionProperties `json:"label_option_properties,omitempty" xml:"label_option_properties,omitempty"`
	// 是否可以绑定多个定向,获取不到 默认为否,true:是 false:否
	IsMulti bool `json:"is_multi,omitempty" xml:"is_multi,omitempty"`
}
