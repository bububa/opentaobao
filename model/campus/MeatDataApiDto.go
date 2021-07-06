package campus

// MeatDataApiDto 结构体
type MeatDataApiDto struct {
	// 参数code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 地图映射
	MappingPoint string `json:"mapping_point,omitempty" xml:"mapping_point,omitempty"`
	// 数据类型
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 被关联设备uuid
	RefDeviceId string `json:"ref_device_id,omitempty" xml:"ref_device_id,omitempty"`
	// 被关联属性code
	RefPropertyCode string `json:"ref_property_code,omitempty" xml:"ref_property_code,omitempty"`
	// 是否报警值
	Alarm bool `json:"alarm,omitempty" xml:"alarm,omitempty"`
}
