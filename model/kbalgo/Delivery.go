package kbalgo

import (
	"sync"
)

// Delivery 结构体
type Delivery struct {
	// 分
	Min string `json:"min,omitempty" xml:"min,omitempty"`
	// step_min
	StepMin string `json:"step_min,omitempty" xml:"step_min,omitempty"`
	// step_base
	StepBase string `json:"step_base,omitempty" xml:"step_base,omitempty"`
}

var poolDelivery = sync.Pool{
	New: func() any {
		return new(Delivery)
	},
}

// GetDelivery() 从对象池中获取Delivery
func GetDelivery() *Delivery {
	return poolDelivery.Get().(*Delivery)
}

// ReleaseDelivery 释放Delivery
func ReleaseDelivery(v *Delivery) {
	v.Min = ""
	v.StepMin = ""
	v.StepBase = ""
	poolDelivery.Put(v)
}
