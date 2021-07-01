package category

// PropsModificationResult 结构体
type PropsModificationResult struct {
	// 类目属性Id
	PropertyId int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 是否必填，0-否，1-是
	Required int64 `json:"required,omitempty" xml:"required,omitempty"`
	// 是否多选，0-否，1-是
	MultiSelect int64 `json:"multi_select,omitempty" xml:"multi_select,omitempty"`
	// 变更类型: 删除(1), 修改(2), 新增(3)
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 变更日期
	Ds string `json:"ds,omitempty" xml:"ds,omitempty"`
	// 属性名称
	PropertyName string `json:"property_name,omitempty" xml:"property_name,omitempty"`
}
