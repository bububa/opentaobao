package campus

// GuardConfigDto 结构体
type GuardConfigDto struct {
	// 设备
	InputList []SubDeviceDto `json:"input_list,omitempty" xml:"input_list>sub_device_dto,omitempty"`
	// 子设备
	OutputList []SubDeviceDto `json:"output_list,omitempty" xml:"output_list>sub_device_dto,omitempty"`
	// guard
	Guard *Guard `json:"guard,omitempty" xml:"guard,omitempty"`
}
