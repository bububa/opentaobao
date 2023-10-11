package simba

// LabelConfigVo 结构体
type LabelConfigVo struct {
	// 标签tab名称
	LabelGroupName string `json:"label_group_name,omitempty" xml:"label_group_name,omitempty"`
	// 标签业务类型，extend:扩展人群,seed:种子人群
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 标签id
	LabelId int64 `json:"label_id,omitempty" xml:"label_id,omitempty"`
}
