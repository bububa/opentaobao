package tmallgenie

import (
	"sync"
)

// ScheduleDto 结构体
type ScheduleDto struct {
	// 调度类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 一次性
	Once *OnceSchedule `json:"once,omitempty" xml:"once,omitempty"`
	// 法定工作日
	StatutoryWorkingDay *StatutoryWorkingDaySchedule `json:"statutory_working_day,omitempty" xml:"statutory_working_day,omitempty"`
	// 每周
	Weekly *WeeklySchedule `json:"weekly,omitempty" xml:"weekly,omitempty"`
}

var poolScheduleDto = sync.Pool{
	New: func() any {
		return new(ScheduleDto)
	},
}

// GetScheduleDto() 从对象池中获取ScheduleDto
func GetScheduleDto() *ScheduleDto {
	return poolScheduleDto.Get().(*ScheduleDto)
}

// ReleaseScheduleDto 释放ScheduleDto
func ReleaseScheduleDto(v *ScheduleDto) {
	v.Type = ""
	v.Once = nil
	v.StatutoryWorkingDay = nil
	v.Weekly = nil
	poolScheduleDto.Put(v)
}
