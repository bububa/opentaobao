package drugtrace

// BaseInfosDTO 
type BaseInfosDTO struct {
    // 药品基础信息
    BaseInfoList   []BaseInfoDTO `json:"base_info_list,omitempty" xml:"base_info_list>base_info_dto,omitempty"`
}
