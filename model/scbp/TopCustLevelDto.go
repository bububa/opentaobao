package scbp

import (
	"sync"
)

// TopCustLevelDto 结构体
type TopCustLevelDto struct {
	// growthLevel
	GrowthLevel string `json:"growth_level,omitempty" xml:"growth_level,omitempty"`
	// levelScore
	LevelScore int64 `json:"level_score,omitempty" xml:"level_score,omitempty"`
}

var poolTopCustLevelDto = sync.Pool{
	New: func() any {
		return new(TopCustLevelDto)
	},
}

// GetTopCustLevelDto() 从对象池中获取TopCustLevelDto
func GetTopCustLevelDto() *TopCustLevelDto {
	return poolTopCustLevelDto.Get().(*TopCustLevelDto)
}

// ReleaseTopCustLevelDto 释放TopCustLevelDto
func ReleaseTopCustLevelDto(v *TopCustLevelDto) {
	v.GrowthLevel = ""
	v.LevelScore = 0
	poolTopCustLevelDto.Put(v)
}
