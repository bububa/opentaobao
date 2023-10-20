package drugtrace

import (
	"sync"
)

// ExtInfoDto 结构体
type ExtInfoDto struct {
	// 生产信息模块
	ProduceInfoDto *ProduceInfoDto `json:"produce_info_dto,omitempty" xml:"produce_info_dto,omitempty"`
	// 种源与基地信息
	SeedlingsBaseInfoDto *SeedlingsBaseInfoDto `json:"seedlings_base_info_dto,omitempty" xml:"seedlings_base_info_dto,omitempty"`
	// 饮片检验信息
	PiecesDetectionDto *PiecesDetectionDto `json:"pieces_detection_dto,omitempty" xml:"pieces_detection_dto,omitempty"`
	// 药材检验信息
	MaterialsDetectionDto *MaterialsDetectionDto `json:"materials_detection_dto,omitempty" xml:"materials_detection_dto,omitempty"`
	// 采购管理
	PurchaseInfoDto *PurchaseInfoDto `json:"purchase_info_dto,omitempty" xml:"purchase_info_dto,omitempty"`
	// 产地初加工管理
	ProcessInfoDto *ProcessInfoDto `json:"process_info_dto,omitempty" xml:"process_info_dto,omitempty"`
	// 生产管理
	PiecesProduceInfoDto *PiecesProduceInfoDto `json:"pieces_produce_info_dto,omitempty" xml:"pieces_produce_info_dto,omitempty"`
	// 存储管理
	WarehouseInfoDto *WarehouseInfoDto `json:"warehouse_info_dto,omitempty" xml:"warehouse_info_dto,omitempty"`
	// 种植管理
	PlantingInfoDto *PlantingInfoDto `json:"planting_info_dto,omitempty" xml:"planting_info_dto,omitempty"`
	// 企业信息
	EntInfoDto *EntInfoDto `json:"ent_info_dto,omitempty" xml:"ent_info_dto,omitempty"`
}

var poolExtInfoDto = sync.Pool{
	New: func() any {
		return new(ExtInfoDto)
	},
}

// GetExtInfoDto() 从对象池中获取ExtInfoDto
func GetExtInfoDto() *ExtInfoDto {
	return poolExtInfoDto.Get().(*ExtInfoDto)
}

// ReleaseExtInfoDto 释放ExtInfoDto
func ReleaseExtInfoDto(v *ExtInfoDto) {
	v.ProduceInfoDto = nil
	v.SeedlingsBaseInfoDto = nil
	v.PiecesDetectionDto = nil
	v.MaterialsDetectionDto = nil
	v.PurchaseInfoDto = nil
	v.ProcessInfoDto = nil
	v.PiecesProduceInfoDto = nil
	v.WarehouseInfoDto = nil
	v.PlantingInfoDto = nil
	v.EntInfoDto = nil
	poolExtInfoDto.Put(v)
}
