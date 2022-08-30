package tmallsc

// ServiceDefinition 结构体
type ServiceDefinition struct {
	// 服务内容list
	ServiceContents []ServiceContent `json:"service_contents,omitempty" xml:"service_contents>service_content,omitempty"`
	// 服务属性
	ServiceProperties []ServiceProperty `json:"service_properties,omitempty" xml:"service_properties>service_property,omitempty"`
	// 服务对象类型
	ServiceObjectType string `json:"service_object_type,omitempty" xml:"service_object_type,omitempty"`
	// 服务对象类型名称
	ServiceObjectTypeName string `json:"service_object_type_name,omitempty" xml:"service_object_type_name,omitempty"`
}
