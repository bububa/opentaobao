package drugtrace

import (
	"sync"
)

// DrugScancodeFlowLogDto 结构体
type DrugScancodeFlowLogDto struct {
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 企业名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolDrugScancodeFlowLogDto = sync.Pool{
	New: func() any {
		return new(DrugScancodeFlowLogDto)
	},
}

// GetDrugScancodeFlowLogDto() 从对象池中获取DrugScancodeFlowLogDto
func GetDrugScancodeFlowLogDto() *DrugScancodeFlowLogDto {
	return poolDrugScancodeFlowLogDto.Get().(*DrugScancodeFlowLogDto)
}

// ReleaseDrugScancodeFlowLogDto 释放DrugScancodeFlowLogDto
func ReleaseDrugScancodeFlowLogDto(v *DrugScancodeFlowLogDto) {
	v.Code = ""
	v.Name = ""
	poolDrugScancodeFlowLogDto.Put(v)
}
