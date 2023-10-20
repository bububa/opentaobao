package alidoc

import (
	"sync"
)

// DrugUsageDto 结构体
type DrugUsageDto struct {
	// 用法用量单位
	MeasureUnit string `json:"measure_unit,omitempty" xml:"measure_unit,omitempty"`
	// 用法
	DrugUsage string `json:"drug_usage,omitempty" xml:"drug_usage,omitempty"`
	// 天数
	Days string `json:"days,omitempty" xml:"days,omitempty"`
	// 频次值单位
	FrequencyUnit string `json:"frequency_unit,omitempty" xml:"frequency_unit,omitempty"`
	// 频次值
	FrequencyValue string `json:"frequency_value,omitempty" xml:"frequency_value,omitempty"`
	// 频次
	Frequency string `json:"frequency,omitempty" xml:"frequency,omitempty"`
	// 每次用量
	DoseValue string `json:"dose_value,omitempty" xml:"dose_value,omitempty"`
}

var poolDrugUsageDto = sync.Pool{
	New: func() any {
		return new(DrugUsageDto)
	},
}

// GetDrugUsageDto() 从对象池中获取DrugUsageDto
func GetDrugUsageDto() *DrugUsageDto {
	return poolDrugUsageDto.Get().(*DrugUsageDto)
}

// ReleaseDrugUsageDto 释放DrugUsageDto
func ReleaseDrugUsageDto(v *DrugUsageDto) {
	v.MeasureUnit = ""
	v.DrugUsage = ""
	v.Days = ""
	v.FrequencyUnit = ""
	v.FrequencyValue = ""
	v.Frequency = ""
	v.DoseValue = ""
	poolDrugUsageDto.Put(v)
}
