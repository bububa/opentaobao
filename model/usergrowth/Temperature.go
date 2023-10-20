package usergrowth

import (
	"sync"
)

// Temperature 结构体
type Temperature struct {
	// 当日最高温度
	Highest string `json:"highest,omitempty" xml:"highest,omitempty"`
	// 当日最低温度
	Lowest string `json:"lowest,omitempty" xml:"lowest,omitempty"`
}

var poolTemperature = sync.Pool{
	New: func() any {
		return new(Temperature)
	},
}

// GetTemperature() 从对象池中获取Temperature
func GetTemperature() *Temperature {
	return poolTemperature.Get().(*Temperature)
}

// ReleaseTemperature 释放Temperature
func ReleaseTemperature(v *Temperature) {
	v.Highest = ""
	v.Lowest = ""
	poolTemperature.Put(v)
}
