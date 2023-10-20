package campus

import (
	"sync"
)

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

var poolGuardConfigDto = sync.Pool{
	New: func() any {
		return new(GuardConfigDto)
	},
}

// GetGuardConfigDto() 从对象池中获取GuardConfigDto
func GetGuardConfigDto() *GuardConfigDto {
	return poolGuardConfigDto.Get().(*GuardConfigDto)
}

// ReleaseGuardConfigDto 释放GuardConfigDto
func ReleaseGuardConfigDto(v *GuardConfigDto) {
	v.InputList = v.InputList[:0]
	v.OutputList = v.OutputList[:0]
	v.Guard = nil
	v.OpenPlanId = 0
	v.BlockPlanId = 0
	poolGuardConfigDto.Put(v)
}
