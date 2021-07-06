package feedflow

// LabelDto 结构体
type LabelDto struct {
	// 选项结构
	Options []OptionDto `json:"options,omitempty" xml:"options>option_dto,omitempty"`
	// 标签值，可通过标签接口获取
	LabelValue string `json:"label_value,omitempty" xml:"label_value,omitempty"`
	// 定向类型
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 标签名称
	LabelName string `json:"label_name,omitempty" xml:"label_name,omitempty"`
	// 标签描述
	LabelDesc string `json:"label_desc,omitempty" xml:"label_desc,omitempty"`
	// 标签id，可通过标签接口获取
	LabelId int64 `json:"label_id,omitempty" xml:"label_id,omitempty"`
	// 定向id，可通过标签接口获取
	TargetId int64 `json:"target_id,omitempty" xml:"target_id,omitempty"`
}
