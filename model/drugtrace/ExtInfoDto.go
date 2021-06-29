package drugtrace

// ExtInfoDTO 
type ExtInfoDTO struct {
    // 生产信息模块
    ProduceInfoDto   *ProduceInfoDTO `json:"produce_info_dto,omitempty" xml:"produce_info_dto,omitempty"`
    // 种源与基地信息
    SeedlingsBaseInfoDto   *SeedlingsBaseInfoDTO `json:"seedlings_base_info_dto,omitempty" xml:"seedlings_base_info_dto,omitempty"`
    // 饮片检验信息
    PiecesDetectionDto   *PiecesDetectionDTO `json:"pieces_detection_dto,omitempty" xml:"pieces_detection_dto,omitempty"`
    // 药材检验信息
    MaterialsDetectionDto   *MaterialsDetectionDTO `json:"materials_detection_dto,omitempty" xml:"materials_detection_dto,omitempty"`
    // 采购管理
    PurchaseInfoDto   *PurchaseInfoDTO `json:"purchase_info_dto,omitempty" xml:"purchase_info_dto,omitempty"`
    // 产地初加工管理
    ProcessInfoDto   *ProcessInfoDTO `json:"process_info_dto,omitempty" xml:"process_info_dto,omitempty"`
    // 生产管理
    PiecesProduceInfoDto   *PiecesProduceInfoDTO `json:"pieces_produce_info_dto,omitempty" xml:"pieces_produce_info_dto,omitempty"`
    // 存储管理
    WarehouseInfoDto   *WarehouseInfoDTO `json:"warehouse_info_dto,omitempty" xml:"warehouse_info_dto,omitempty"`
    // 种植管理
    PlantingInfoDto   *PlantingInfoDTO `json:"planting_info_dto,omitempty" xml:"planting_info_dto,omitempty"`
    // 企业信息
    EntInfoDto   *EntInfoDTO `json:"ent_info_dto,omitempty" xml:"ent_info_dto,omitempty"`
}
