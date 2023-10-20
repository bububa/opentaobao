package bus

import (
	"sync"
)

// BusSeatPriceRq 结构体
type BusSeatPriceRq struct {
	// 车次ID
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
}

var poolBusSeatPriceRq = sync.Pool{
	New: func() any {
		return new(BusSeatPriceRq)
	},
}

// GetBusSeatPriceRq() 从对象池中获取BusSeatPriceRq
func GetBusSeatPriceRq() *BusSeatPriceRq {
	return poolBusSeatPriceRq.Get().(*BusSeatPriceRq)
}

// ReleaseBusSeatPriceRq 释放BusSeatPriceRq
func ReleaseBusSeatPriceRq(v *BusSeatPriceRq) {
	v.ScheduleId = ""
	poolBusSeatPriceRq.Put(v)
}
