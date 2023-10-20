package axindata

import (
	"sync"
)

// BreakfastDto 结构体
type BreakfastDto struct {
	// 是否含早餐
	Breakfast string `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	// 早餐数量
	BreakfastCount int64 `json:"breakfast_count,omitempty" xml:"breakfast_count,omitempty"`
	// 规则生效开始时间
	StartDate int64 `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 规则生效结束时间(不包含)
	EndDate int64 `json:"end_date,omitempty" xml:"end_date,omitempty"`
}

var poolBreakfastDto = sync.Pool{
	New: func() any {
		return new(BreakfastDto)
	},
}

// GetBreakfastDto() 从对象池中获取BreakfastDto
func GetBreakfastDto() *BreakfastDto {
	return poolBreakfastDto.Get().(*BreakfastDto)
}

// ReleaseBreakfastDto 释放BreakfastDto
func ReleaseBreakfastDto(v *BreakfastDto) {
	v.Breakfast = ""
	v.BreakfastCount = 0
	v.StartDate = 0
	v.EndDate = 0
	poolBreakfastDto.Put(v)
}
