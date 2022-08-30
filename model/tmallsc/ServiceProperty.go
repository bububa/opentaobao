package tmallsc

// ServiceProperty 结构体
type ServiceProperty struct {
	// 服务属性值枚举
	ServicePropertyValueEnum []string `json:"service_property_value_enum,omitempty" xml:"service_property_value_enum>string,omitempty"`
	// 服务属性名称
	ServicePropertyName string `json:"service_property_name,omitempty" xml:"service_property_name,omitempty"`
}
