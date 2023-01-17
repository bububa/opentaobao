package flight

// ProcessingDto 结构体
type ProcessingDto struct {
	// 附件列表
	FileInfoDtoList []UploadFileInfoDto `json:"file_info_dto_list,omitempty" xml:"file_info_dto_list>upload_file_info_dto,omitempty"`
	// 协同回复内容
	Reply string `json:"reply,omitempty" xml:"reply,omitempty"`
	// 协同单ID
	CaseId int64 `json:"case_id,omitempty" xml:"case_id,omitempty"`
	// 1:国内，2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
	// 协同单extraInfo总
	TotalCaseBaseExtraInfoDto *TotalCaseExtraInfoDto `json:"total_case_base_extra_info_dto,omitempty" xml:"total_case_base_extra_info_dto,omitempty"`
}
