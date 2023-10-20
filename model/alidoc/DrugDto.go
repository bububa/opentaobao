package alidoc

import (
	"sync"
)

// DrugDto 结构体
type DrugDto struct {
	// 药品用法用量
	DrugUsageList []DrugUsageDto `json:"drug_usage_list,omitempty" xml:"drug_usage_list>drug_usage_dto,omitempty"`
	// 药品Id
	DrugId string `json:"drug_id,omitempty" xml:"drug_id,omitempty"`
	// 药品名称
	DrugName string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
	// 规格
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
	// 数量
	Total string `json:"total,omitempty" xml:"total,omitempty"`
	// 剂型
	DoseFrom string `json:"dose_from,omitempty" xml:"dose_from,omitempty"`
}

var poolDrugDto = sync.Pool{
	New: func() any {
		return new(DrugDto)
	},
}

// GetDrugDto() 从对象池中获取DrugDto
func GetDrugDto() *DrugDto {
	return poolDrugDto.Get().(*DrugDto)
}

// ReleaseDrugDto 释放DrugDto
func ReleaseDrugDto(v *DrugDto) {
	v.DrugUsageList = v.DrugUsageList[:0]
	v.DrugId = ""
	v.DrugName = ""
	v.Spec = ""
	v.Total = ""
	v.DoseFrom = ""
	poolDrugDto.Put(v)
}
