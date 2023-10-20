package scbp

import (
	"sync"
)

// TargetReportDto 结构体
type TargetReportDto struct {
	// 返回实体集合
	TargetEffectList []TargetEffectDto `json:"target_effect_list,omitempty" xml:"target_effect_list>target_effect_dto,omitempty"`
}

var poolTargetReportDto = sync.Pool{
	New: func() any {
		return new(TargetReportDto)
	},
}

// GetTargetReportDto() 从对象池中获取TargetReportDto
func GetTargetReportDto() *TargetReportDto {
	return poolTargetReportDto.Get().(*TargetReportDto)
}

// ReleaseTargetReportDto 释放TargetReportDto
func ReleaseTargetReportDto(v *TargetReportDto) {
	v.TargetEffectList = v.TargetEffectList[:0]
	poolTargetReportDto.Put(v)
}
