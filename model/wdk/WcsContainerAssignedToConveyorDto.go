package wdk

// WcsContainerAssignedToConveyorDto 结构体
type WcsContainerAssignedToConveyorDto struct {
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// warehouseId
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// batchCode
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// containerCode
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// waveCode
	WaveCode string `json:"wave_code,omitempty" xml:"wave_code,omitempty"`
	// conveyorId
	ConveyorId int64 `json:"conveyor_id,omitempty" xml:"conveyor_id,omitempty"`
	// conveyorAssignedTime
	ConveyorAssignedTime string `json:"conveyor_assigned_time,omitempty" xml:"conveyor_assigned_time,omitempty"`
	// assignedConveyor
	AssignedConveyor int64 `json:"assigned_conveyor,omitempty" xml:"assigned_conveyor,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}
