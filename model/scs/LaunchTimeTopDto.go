package scs

import (
	"sync"
)

// LaunchTimeTopDto 结构体
type LaunchTimeTopDto struct {
	// 计划开始时间,需要为0点的时间，大于今天
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 计划结束时间,需要为0点的时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 是否永远持续，持续推广为true
	LaunchForever bool `json:"launch_forever,omitempty" xml:"launch_forever,omitempty"`
}

var poolLaunchTimeTopDto = sync.Pool{
	New: func() any {
		return new(LaunchTimeTopDto)
	},
}

// GetLaunchTimeTopDto() 从对象池中获取LaunchTimeTopDto
func GetLaunchTimeTopDto() *LaunchTimeTopDto {
	return poolLaunchTimeTopDto.Get().(*LaunchTimeTopDto)
}

// ReleaseLaunchTimeTopDto 释放LaunchTimeTopDto
func ReleaseLaunchTimeTopDto(v *LaunchTimeTopDto) {
	v.BeginTime = ""
	v.EndTime = ""
	v.LaunchForever = false
	poolLaunchTimeTopDto.Put(v)
}
