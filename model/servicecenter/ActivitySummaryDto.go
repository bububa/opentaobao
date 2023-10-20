package servicecenter

import (
	"sync"
)

// ActivitySummaryDto 结构体
type ActivitySummaryDto struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间（设置）
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 中止时间
	TerminateTime string `json:"terminate_time,omitempty" xml:"terminate_time,omitempty"`
}

var poolActivitySummaryDto = sync.Pool{
	New: func() any {
		return new(ActivitySummaryDto)
	},
}

// GetActivitySummaryDto() 从对象池中获取ActivitySummaryDto
func GetActivitySummaryDto() *ActivitySummaryDto {
	return poolActivitySummaryDto.Get().(*ActivitySummaryDto)
}

// ReleaseActivitySummaryDto 释放ActivitySummaryDto
func ReleaseActivitySummaryDto(v *ActivitySummaryDto) {
	v.StartTime = ""
	v.EndTime = ""
	v.TerminateTime = ""
	poolActivitySummaryDto.Put(v)
}
