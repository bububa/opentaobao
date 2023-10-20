package alidoc

import (
	"sync"
)

// DrugDiseaseMappingDto 结构体
type DrugDiseaseMappingDto struct {
	// 诊断信息
	DiseaseList []DiseaseInfo `json:"disease_list,omitempty" xml:"disease_list>disease_info,omitempty"`
	// spuid
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}

var poolDrugDiseaseMappingDto = sync.Pool{
	New: func() any {
		return new(DrugDiseaseMappingDto)
	},
}

// GetDrugDiseaseMappingDto() 从对象池中获取DrugDiseaseMappingDto
func GetDrugDiseaseMappingDto() *DrugDiseaseMappingDto {
	return poolDrugDiseaseMappingDto.Get().(*DrugDiseaseMappingDto)
}

// ReleaseDrugDiseaseMappingDto 释放DrugDiseaseMappingDto
func ReleaseDrugDiseaseMappingDto(v *DrugDiseaseMappingDto) {
	v.DiseaseList = v.DiseaseList[:0]
	v.SpuId = ""
	poolDrugDiseaseMappingDto.Put(v)
}
