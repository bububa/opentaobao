package icbudropshipping

import (
	"sync"
)

// LadderPeriod 结构体
type LadderPeriod struct {
	// max quantity
	MaxQuantity int64 `json:"max_quantity,omitempty" xml:"max_quantity,omitempty"`
	// min quantity
	MinQuantity int64 `json:"min_quantity,omitempty" xml:"min_quantity,omitempty"`
	// Delivery time
	ProcessPeriod int64 `json:"process_period,omitempty" xml:"process_period,omitempty"`
}

var poolLadderPeriod = sync.Pool{
	New: func() any {
		return new(LadderPeriod)
	},
}

// GetLadderPeriod() 从对象池中获取LadderPeriod
func GetLadderPeriod() *LadderPeriod {
	return poolLadderPeriod.Get().(*LadderPeriod)
}

// ReleaseLadderPeriod 释放LadderPeriod
func ReleaseLadderPeriod(v *LadderPeriod) {
	v.MaxQuantity = 0
	v.MinQuantity = 0
	v.ProcessPeriod = 0
	poolLadderPeriod.Put(v)
}
