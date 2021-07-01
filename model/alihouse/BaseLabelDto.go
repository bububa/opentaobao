package alihouse

// BaseLabelDto 结构体
type BaseLabelDto struct {
	// 父级业务
	ParentBusiness string `json:"parent_business,omitempty" xml:"parent_business,omitempty"`
	// 业务
	Business string `json:"business,omitempty" xml:"business,omitempty"`
	// 分类
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 标签ID
	LabelId string `json:"label_id,omitempty" xml:"label_id,omitempty"`
	// 标签名称
	LabelName string `json:"label_name,omitempty" xml:"label_name,omitempty"`
	// 类型ID
	TypeId string `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// 业务ID
	BusinessId string `json:"business_id,omitempty" xml:"business_id,omitempty"`
	// 父级业务ID
	ParentBusinessId string `json:"parent_business_id,omitempty" xml:"parent_business_id,omitempty"`
}
