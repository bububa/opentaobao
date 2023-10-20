package drugtrace

import (
	"sync"
)

// CodeQueryDrugInfoDto 结构体
type CodeQueryDrugInfoDto struct {
	// 药品信息列表
	InfoList []CodeFullInfoDto `json:"info_list,omitempty" xml:"info_list>code_full_info_dto,omitempty"`
}

var poolCodeQueryDrugInfoDto = sync.Pool{
	New: func() any {
		return new(CodeQueryDrugInfoDto)
	},
}

// GetCodeQueryDrugInfoDto() 从对象池中获取CodeQueryDrugInfoDto
func GetCodeQueryDrugInfoDto() *CodeQueryDrugInfoDto {
	return poolCodeQueryDrugInfoDto.Get().(*CodeQueryDrugInfoDto)
}

// ReleaseCodeQueryDrugInfoDto 释放CodeQueryDrugInfoDto
func ReleaseCodeQueryDrugInfoDto(v *CodeQueryDrugInfoDto) {
	v.InfoList = v.InfoList[:0]
	poolCodeQueryDrugInfoDto.Put(v)
}
