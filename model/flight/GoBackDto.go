package flight

// GoBackDto 结构体
type GoBackDto struct {
	// 附件列表
	FileInfoDtoList []UploadFileInfoDto `json:"file_info_dto_list,omitempty" xml:"file_info_dto_list>upload_file_info_dto,omitempty"`
	// 备注信息
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 协同单Id
	CaseId int64 `json:"case_id,omitempty" xml:"case_id,omitempty"`
}
