package flight

import (
	"sync"
)

// RejectDto 结构体
type RejectDto struct {
	// 附件列表
	FileInfoDtoList []UploadFileInfoDto `json:"file_info_dto_list,omitempty" xml:"file_info_dto_list>upload_file_info_dto,omitempty"`
	// 拒绝原因
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 协同单id
	CaseId int64 `json:"case_id,omitempty" xml:"case_id,omitempty"`
	// 1:国内，2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolRejectDto = sync.Pool{
	New: func() any {
		return new(RejectDto)
	},
}

// GetRejectDto() 从对象池中获取RejectDto
func GetRejectDto() *RejectDto {
	return poolRejectDto.Get().(*RejectDto)
}

// ReleaseRejectDto 释放RejectDto
func ReleaseRejectDto(v *RejectDto) {
	v.FileInfoDtoList = v.FileInfoDtoList[:0]
	v.Content = ""
	v.CaseId = 0
	v.DomesticIntl = 0
	poolRejectDto.Put(v)
}
