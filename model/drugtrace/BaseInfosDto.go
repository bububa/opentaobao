package drugtrace

// BaseInfosDto 
type BaseInfosDto struct {
    // 药品基础信息
    BaseInfoList   []BaseInfoDto `json:"base_info_list,omitempty" xml:"base_info_list>base_info_dto,omitempty"`
}
