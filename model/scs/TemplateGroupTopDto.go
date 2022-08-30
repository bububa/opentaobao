package scs

// TemplateGroupTopDto 结构体
type TemplateGroupTopDto struct {
	// 人群名称
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// 人群描述
	GroupDesc string `json:"group_desc,omitempty" xml:"group_desc,omitempty"`
	// 人群id
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
}
