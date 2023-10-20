package wdk

import (
	"sync"
)

// WcsContainerAssignedToConveyorDto 结构体
type WcsContainerAssignedToConveyorDto struct {
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// batchCode
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// containerCode
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// waveCode
	WaveCode string `json:"wave_code,omitempty" xml:"wave_code,omitempty"`
	// conveyorAssignedTime
	ConveyorAssignedTime string `json:"conveyor_assigned_time,omitempty" xml:"conveyor_assigned_time,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// warehouseId
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// conveyorId
	ConveyorId int64 `json:"conveyor_id,omitempty" xml:"conveyor_id,omitempty"`
	// assignedConveyor
	AssignedConveyor int64 `json:"assigned_conveyor,omitempty" xml:"assigned_conveyor,omitempty"`
}

var poolWcsContainerAssignedToConveyorDto = sync.Pool{
	New: func() any {
		return new(WcsContainerAssignedToConveyorDto)
	},
}

// GetWcsContainerAssignedToConveyorDto() 从对象池中获取WcsContainerAssignedToConveyorDto
func GetWcsContainerAssignedToConveyorDto() *WcsContainerAssignedToConveyorDto {
	return poolWcsContainerAssignedToConveyorDto.Get().(*WcsContainerAssignedToConveyorDto)
}

// ReleaseWcsContainerAssignedToConveyorDto 释放WcsContainerAssignedToConveyorDto
func ReleaseWcsContainerAssignedToConveyorDto(v *WcsContainerAssignedToConveyorDto) {
	v.WarehouseCode = ""
	v.BatchCode = ""
	v.ContainerCode = ""
	v.WaveCode = ""
	v.ConveyorAssignedTime = ""
	v.Uuid = ""
	v.WarehouseId = 0
	v.ConveyorId = 0
	v.AssignedConveyor = 0
	poolWcsContainerAssignedToConveyorDto.Put(v)
}
