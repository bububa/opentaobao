package tmallsc

// SparePartsDetailsSaveRequest 结构体
type SparePartsDetailsSaveRequest struct {
	// 备件申请单信息
	ApplicationInfoDto *ApplicationInfoDto `json:"application_info_dto,omitempty" xml:"application_info_dto,omitempty"`
	// 备件明细
	SparePartsInfoDto *SparePartsInfoDto `json:"spare_parts_info_dto,omitempty" xml:"spare_parts_info_dto,omitempty"`
	// 天猫服务工单号
	WorkCardId int64 `json:"work_card_id,omitempty" xml:"work_card_id,omitempty"`
}
