package scs

import (
	"sync"
)

// DayBudgetTopDto 结构体
type DayBudgetTopDto struct {
	// 持续推广计划日预算
	DayBudget string `json:"day_budget,omitempty" xml:"day_budget,omitempty"`
}

var poolDayBudgetTopDto = sync.Pool{
	New: func() any {
		return new(DayBudgetTopDto)
	},
}

// GetDayBudgetTopDto() 从对象池中获取DayBudgetTopDto
func GetDayBudgetTopDto() *DayBudgetTopDto {
	return poolDayBudgetTopDto.Get().(*DayBudgetTopDto)
}

// ReleaseDayBudgetTopDto 释放DayBudgetTopDto
func ReleaseDayBudgetTopDto(v *DayBudgetTopDto) {
	v.DayBudget = ""
	poolDayBudgetTopDto.Put(v)
}
