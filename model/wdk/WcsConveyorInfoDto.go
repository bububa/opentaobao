package wdk

import (
	"sync"
)

// WcsConveyorInfoDto 结构体
type WcsConveyorInfoDto struct {
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// lastHeartbeatTime
	LastHeartbeatTime string `json:"last_heartbeat_time,omitempty" xml:"last_heartbeat_time,omitempty"`
	// warehouseId
	WarehouseId int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// conveyorId
	ConveyorId int64 `json:"conveyor_id,omitempty" xml:"conveyor_id,omitempty"`
	// isRunning
	IsRunning int64 `json:"is_running,omitempty" xml:"is_running,omitempty"`
	// slidewaysUnused
	SlidewaysUnused int64 `json:"slideways_unused,omitempty" xml:"slideways_unused,omitempty"`
	// circlingContainers
	CirclingContainers int64 `json:"circling_containers,omitempty" xml:"circling_containers,omitempty"`
}

var poolWcsConveyorInfoDto = sync.Pool{
	New: func() any {
		return new(WcsConveyorInfoDto)
	},
}

// GetWcsConveyorInfoDto() 从对象池中获取WcsConveyorInfoDto
func GetWcsConveyorInfoDto() *WcsConveyorInfoDto {
	return poolWcsConveyorInfoDto.Get().(*WcsConveyorInfoDto)
}

// ReleaseWcsConveyorInfoDto 释放WcsConveyorInfoDto
func ReleaseWcsConveyorInfoDto(v *WcsConveyorInfoDto) {
	v.WarehouseCode = ""
	v.LastHeartbeatTime = ""
	v.WarehouseId = 0
	v.ConveyorId = 0
	v.IsRunning = 0
	v.SlidewaysUnused = 0
	v.CirclingContainers = 0
	poolWcsConveyorInfoDto.Put(v)
}
