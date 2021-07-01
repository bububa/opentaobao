package drugtrace

// BaseInfosDto 结构体
type BaseInfosDto struct {
	// 药品基础信息
	BaseInfoList []BaseInfoDto `json:"base_info_list,omitempty" xml:"base_info_list>base_info_dto,omitempty"`
}
