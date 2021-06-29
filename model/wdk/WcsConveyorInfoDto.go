package wdk

// WcsConveyorInfoDto 
type WcsConveyorInfoDto struct {
    // warehouseCode
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // warehouseId
    WarehouseId   int64 `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
    // conveyorId
    ConveyorId   int64 `json:"conveyor_id,omitempty" xml:"conveyor_id,omitempty"`
    // isRunning
    IsRunning   int64 `json:"is_running,omitempty" xml:"is_running,omitempty"`
    // slidewaysUnused
    SlidewaysUnused   int64 `json:"slideways_unused,omitempty" xml:"slideways_unused,omitempty"`
    // lastHeartbeatTime
    LastHeartbeatTime   string `json:"last_heartbeat_time,omitempty" xml:"last_heartbeat_time,omitempty"`
    // circlingContainers
    CirclingContainers   int64 `json:"circling_containers,omitempty" xml:"circling_containers,omitempty"`
}
