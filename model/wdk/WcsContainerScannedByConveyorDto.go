package wdk

import (
	"sync"
)

// WcsContainerScannedByConveyorDto 结构体
type WcsContainerScannedByConveyorDto struct {
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// warehouseId
	WarehouseId string `json:"warehouse_id,omitempty" xml:"warehouse_id,omitempty"`
	// batchCode
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// containerCode
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// waveCode
	WaveCode string `json:"wave_code,omitempty" xml:"wave_code,omitempty"`
	// conveyorId
	ConveyorId string `json:"conveyor_id,omitempty" xml:"conveyor_id,omitempty"`
	// scannedTime
	ScannedTime string `json:"scanned_time,omitempty" xml:"scanned_time,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolWcsContainerScannedByConveyorDto = sync.Pool{
	New: func() any {
		return new(WcsContainerScannedByConveyorDto)
	},
}

// GetWcsContainerScannedByConveyorDto() 从对象池中获取WcsContainerScannedByConveyorDto
func GetWcsContainerScannedByConveyorDto() *WcsContainerScannedByConveyorDto {
	return poolWcsContainerScannedByConveyorDto.Get().(*WcsContainerScannedByConveyorDto)
}

// ReleaseWcsContainerScannedByConveyorDto 释放WcsContainerScannedByConveyorDto
func ReleaseWcsContainerScannedByConveyorDto(v *WcsContainerScannedByConveyorDto) {
	v.WarehouseCode = ""
	v.WarehouseId = ""
	v.BatchCode = ""
	v.ContainerCode = ""
	v.WaveCode = ""
	v.ConveyorId = ""
	v.ScannedTime = ""
	v.Uuid = ""
	poolWcsContainerScannedByConveyorDto.Put(v)
}
