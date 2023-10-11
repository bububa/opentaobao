package miniapp

// AfterSaleFieldMetaRecord 结构体
type AfterSaleFieldMetaRecord struct {
	// 子结构
	Children []string `json:"children,omitempty" xml:"children>string,omitempty"`
	// 字段名称
	FieldName string `json:"field_name,omitempty" xml:"field_name,omitempty"`
	// 字段描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 组件类型
	ComponentType string `json:"component_type,omitempty" xml:"component_type,omitempty"`
	// 配置信息
	CustomConfig string `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}
