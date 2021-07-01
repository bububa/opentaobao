package moscm

// PropertyDto 结构体
type PropertyDto struct {
	// 属性别名
	PAlias string `json:"p_alias,omitempty" xml:"p_alias,omitempty"`
	// 自定义属性名称
	PCustomName string `json:"p_custom_name,omitempty" xml:"p_custom_name,omitempty"`
	// 对应天猫属性Id
	PId string `json:"p_id,omitempty" xml:"p_id,omitempty"`
	// 对应天猫属性名称
	PName string `json:"p_name,omitempty" xml:"p_name,omitempty"`
	// 属性值别名
	VAlias string `json:"v_alias,omitempty" xml:"v_alias,omitempty"`
	// 自定义value名称（优先展示此字段）
	VCustomName string `json:"v_custom_name,omitempty" xml:"v_custom_name,omitempty"`
	// 对应天猫属性Id
	VId string `json:"v_id,omitempty" xml:"v_id,omitempty"`
	// 对应天猫属性值名称（v_custom_name不传时展示此字段）
	VName string `json:"v_name,omitempty" xml:"v_name,omitempty"`
	// 花色图片
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 顺序
	Order int64 `json:"order,omitempty" xml:"order,omitempty"`
	// 属性类型（sale, nonCritical, critical, tmallItem）
	PropertyType string `json:"property_type,omitempty" xml:"property_type,omitempty"`
}
