package mc

import (
	"sync"
)

// TaskDto 结构体
type TaskDto struct {
	// 投放计划名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 投放计划id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTaskDto = sync.Pool{
	New: func() any {
		return new(TaskDto)
	},
}

// GetTaskDto() 从对象池中获取TaskDto
func GetTaskDto() *TaskDto {
	return poolTaskDto.Get().(*TaskDto)
}

// ReleaseTaskDto 释放TaskDto
func ReleaseTaskDto(v *TaskDto) {
	v.Name = ""
	v.StartTime = ""
	v.EndTime = ""
	v.Id = 0
	poolTaskDto.Put(v)
}
