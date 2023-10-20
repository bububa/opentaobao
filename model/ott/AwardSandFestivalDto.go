package ott

import (
	"sync"
)

// AwardSandFestivalDto 结构体
type AwardSandFestivalDto struct {
	// awardSandFestival
	AwardSandFestival string `json:"award_sand_festival,omitempty" xml:"award_sand_festival,omitempty"`
}

var poolAwardSandFestivalDto = sync.Pool{
	New: func() any {
		return new(AwardSandFestivalDto)
	},
}

// GetAwardSandFestivalDto() 从对象池中获取AwardSandFestivalDto
func GetAwardSandFestivalDto() *AwardSandFestivalDto {
	return poolAwardSandFestivalDto.Get().(*AwardSandFestivalDto)
}

// ReleaseAwardSandFestivalDto 释放AwardSandFestivalDto
func ReleaseAwardSandFestivalDto(v *AwardSandFestivalDto) {
	v.AwardSandFestival = ""
	poolAwardSandFestivalDto.Put(v)
}
