package feedflow

import (
	"sync"
)

// LaunchPeriodDto 结构体
type LaunchPeriodDto struct {
	// 时间
	TimeSpanList []TimeSpanDto `json:"time_span_list,omitempty" xml:"time_span_list>time_span_dto,omitempty"`
}

var poolLaunchPeriodDto = sync.Pool{
	New: func() any {
		return new(LaunchPeriodDto)
	},
}

// GetLaunchPeriodDto() 从对象池中获取LaunchPeriodDto
func GetLaunchPeriodDto() *LaunchPeriodDto {
	return poolLaunchPeriodDto.Get().(*LaunchPeriodDto)
}

// ReleaseLaunchPeriodDto 释放LaunchPeriodDto
func ReleaseLaunchPeriodDto(v *LaunchPeriodDto) {
	v.TimeSpanList = v.TimeSpanList[:0]
	poolLaunchPeriodDto.Put(v)
}
