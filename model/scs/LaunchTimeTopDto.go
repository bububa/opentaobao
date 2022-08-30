package scs

// LaunchTimeTopDto 结构体
type LaunchTimeTopDto struct {
	// 计划开始时间,需要为0点的时间，大于今天
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 计划结束时间,需要为0点的时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 是否永远持续，持续推广为true
	LaunchForever bool `json:"launch_forever,omitempty" xml:"launch_forever,omitempty"`
}
