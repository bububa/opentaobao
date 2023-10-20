package icbu

import (
	"sync"
)

// DeliverPeriod 结构体
type DeliverPeriod struct {
	// 预计需要发货时间
	ProcessPeriod int64 `json:"process_period,omitempty" xml:"process_period,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolDeliverPeriod = sync.Pool{
	New: func() any {
		return new(DeliverPeriod)
	},
}

// GetDeliverPeriod() 从对象池中获取DeliverPeriod
func GetDeliverPeriod() *DeliverPeriod {
	return poolDeliverPeriod.Get().(*DeliverPeriod)
}

// ReleaseDeliverPeriod 释放DeliverPeriod
func ReleaseDeliverPeriod(v *DeliverPeriod) {
	v.ProcessPeriod = 0
	v.Quantity = 0
	poolDeliverPeriod.Put(v)
}
