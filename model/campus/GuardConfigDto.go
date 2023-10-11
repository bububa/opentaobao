package campus

// GuardConfigDto 结构体
type GuardConfigDto struct {
	// 设备
	InputList []SubDeviceDto `json:"input_list,omitempty" xml:"input_list>sub_device_dto,omitempty"`
	// 子设备
	OutputList []SubDeviceDto `json:"output_list,omitempty" xml:"output_list>sub_device_dto,omitempty"`
	// guard
	Guard *Guard `json:"guard,omitempty" xml:"guard,omitempty"`
	// 常开时间计划
	OpenPlanId int64 `json:"open_plan_id,omitempty" xml:"open_plan_id,omitempty"`
	// 封阻时间计划
	BlockPlanId int64 `json:"block_plan_id,omitempty" xml:"block_plan_id,omitempty"`
}
