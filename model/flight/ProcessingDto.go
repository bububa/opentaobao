package flight

import (
	"sync"
)

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

var poolProcessingDto = sync.Pool{
	New: func() any {
		return new(ProcessingDto)
	},
}

// GetProcessingDto() 从对象池中获取ProcessingDto
func GetProcessingDto() *ProcessingDto {
	return poolProcessingDto.Get().(*ProcessingDto)
}

// ReleaseProcessingDto 释放ProcessingDto
func ReleaseProcessingDto(v *ProcessingDto) {
	v.FileInfoDtoList = v.FileInfoDtoList[:0]
	v.Reply = ""
	v.CaseId = 0
	v.DomesticIntl = 0
	v.TotalCaseBaseExtraInfoDto = nil
	poolProcessingDto.Put(v)
}
