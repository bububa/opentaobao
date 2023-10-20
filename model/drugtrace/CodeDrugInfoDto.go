package drugtrace

import (
	"sync"
)

// CodeDrugInfoDto 结构体
type CodeDrugInfoDto struct {
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 码生产信息对象
	CodeProduceInfoDTO *CodeProduceInfoDto `json:"code_produce_info_d_t_o,omitempty" xml:"code_produce_info_d_t_o,omitempty"`
	// 药品基本信息对象
	DrugEntBaseDTO *DrugEntBaseDto `json:"drug_ent_base_d_t_o,omitempty" xml:"drug_ent_base_d_t_o,omitempty"`
}

var poolCodeDrugInfoDto = sync.Pool{
	New: func() any {
		return new(CodeDrugInfoDto)
	},
}

// GetCodeDrugInfoDto() 从对象池中获取CodeDrugInfoDto
func GetCodeDrugInfoDto() *CodeDrugInfoDto {
	return poolCodeDrugInfoDto.Get().(*CodeDrugInfoDto)
}

// ReleaseCodeDrugInfoDto 释放CodeDrugInfoDto
func ReleaseCodeDrugInfoDto(v *CodeDrugInfoDto) {
	v.Code = ""
	v.CodeProduceInfoDTO = nil
	v.DrugEntBaseDTO = nil
	poolCodeDrugInfoDto.Put(v)
}
