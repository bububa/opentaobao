package alilabs

import (
	"sync"
)

// Status 结构体
type Status struct {
	// 开关状态
	Powerstate string `json:"powerstate,omitempty" xml:"powerstate,omitempty"`
	// 模式
	Mode string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 温度
	Temperature string `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// 亮度
	Brightness string `json:"brightness,omitempty" xml:"brightness,omitempty"`
	// 风速
	Windspeed string `json:"windspeed,omitempty" xml:"windspeed,omitempty"`
}

var poolStatus = sync.Pool{
	New: func() any {
		return new(Status)
	},
}

// GetStatus() 从对象池中获取Status
func GetStatus() *Status {
	return poolStatus.Get().(*Status)
}

// ReleaseStatus 释放Status
func ReleaseStatus(v *Status) {
	v.Powerstate = ""
	v.Mode = ""
	v.Temperature = ""
	v.Brightness = ""
	v.Windspeed = ""
	poolStatus.Put(v)
}
